package store

type Users struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	CardNumber string `json:"CardNumber"`
	Balance    int    `json:"Balance"`
	Reserved   int    `json:"Reserved"`
}
