package command

type MovieCommand struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Stars    string `json:"stars"`
}
