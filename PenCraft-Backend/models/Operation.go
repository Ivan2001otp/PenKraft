package models

type Operation struct {
	Operation_type string `json:"operation"`
	Data          Blog   `json:"data"`
}
