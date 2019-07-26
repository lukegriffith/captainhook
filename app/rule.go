package app

type rule struct {
  id string `json:"id"`
  destination_url  string `json:"destination "`
  template string `json:"template"`
  verify_ssl string `json:"verify_ssl"`
}
