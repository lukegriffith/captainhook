package captainhook

type Rule struct {
	Destination_url string `json:"destination"`
	Template        string `json:"template"`
	Verify_ssl      string `json:"verify_ssl"`
}

func (rule Rule) Execute(body string) {

}

