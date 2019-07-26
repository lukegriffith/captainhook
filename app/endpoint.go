package app

type endpoint struct {
	id     string `json:"id"`
	name   string `json:"name"`
	secret string `json:"secret"`
	rules  []rule `json:"rules "`
	source string `json:"source"`
}
