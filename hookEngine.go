package captainhook

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"text/template"
  "github.com/lukemgriffith/captainhook/util"
)

// Main routine that processes received hooks, obtaining endpoints and processing rules.
// Various error checking and validation happens at this stage, i.e mapping required secrets to
// dataBag. data bag is a map of input parameters passed to each rules function.
func Hook(w http.ResponseWriter, r *http.Request, endpoint *Endpoint, secrets SecretEngine, log *util.Logger, dataBag *map[string]interface{}) {

	var bag = *dataBag

	var secretMap = make(map[string]string)

	// Attach secrets to secrets map/
	for _, secret := range endpoint.Secrets {
		v, err := secrets.GetTextSecret(secret)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("unable to get secret from engine", secret)
			return
		}

		secretMap[secret] = v
	}

	// Store secrets on data bag.
	bag["_secrets"] = secretMap

	var request bytes.Buffer
	rules, err := endpoint.GetRules()

	if err != nil {
		log.Println("unable to enumerate rules, endpoint", endpoint.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, r := range rules {

		// For each rule, assign related function then execute via the interface method execute passing in
		// dataBag and the http request.
		AssignFunction(&r)
		err = r.Execute(&request, *dataBag)
		if err != nil {
			log.Println(r, "failed to execute template.", err)
			continue
		}
		// For rules with destinations specified, move to forward function result.
		if r.Destination != "" {
			// Template url and create http request.
			client := &http.Client{}
			dest, err := templateString(r.Destination, dataBag)

			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("forwarding to", dest)
			req, err := http.NewRequest("POST", dest, &request)
			if err != nil {
				log.Println("Failed to create new request", err)
				continue
			}
			// Render headers and attach to request
			for k, v := range r.Headers {
				value, err := templateString(v, dataBag)
				if err != nil {
					log.Println(err)
					continue
				}
				req.Header.Add(k, value)
			}
			// POST request
			resp, err := client.Do(req)
			if err != nil {
				log.Println("post request to", r.Destination, "failed.")
				continue
			}

			if resp.StatusCode > 299 || resp.StatusCode < 200 {
				log.Println("Request returned non 200 status code:", resp.StatusCode)
				log.Println("HTTP Status:", resp.Status)
			}

		}

		request.Reset()
	}
}

// Helper function to quickly render a string template and return the value.
// Useful for URI and Header rendering.
func templateString(templ string, data *map[string]interface{}) (string, error) {

	tmpl, err := template.New("tmpl").Parse(templ)
	if err != nil {
		return "", errors.New(fmt.Sprint("Unable to create template for header from: ", templ))
	}

	buf := make([]byte, 0, 1)
	var tpl *bytes.Buffer = bytes.NewBuffer(buf)
	err = tmpl.Execute(tpl, &data)
	return tpl.String(), nil
}
