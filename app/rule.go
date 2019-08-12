package app

type Rule struct {
	Id              string `json:"id"`
	Destination_url string `json:"destination "`
	Template        string `json:"template"`
	Verify_ssl      string `json:"verify_ssl"`
}
