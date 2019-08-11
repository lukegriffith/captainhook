package app

type endpoint struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
	Rules  []rule `json:"rules"`
	Source string `json:"source"`
}

