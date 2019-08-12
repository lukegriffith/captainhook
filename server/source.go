package server

type Source struct {
	Id         string `json:"id"`
	SourceType string `json:"sourcetype"`
	Location   string `json:"location"`
}
