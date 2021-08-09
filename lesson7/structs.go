package lesson7

//easyjson:json
type News struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Anonce string `json:"anonce"`
	Desc   string `json:"desc"`
}
