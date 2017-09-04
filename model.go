package main

type vehicle struct {
	Plate  string `json:"plate"`
	Model  int    `json:"model"`
	Color  string `json:"color"`
	Type   string `json:"type"`
	IsMain bool   `json:"is_main"`
}

type configs string

type user struct {
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Vehicle  []vehicle `json:"vehicles"`
	Whf      string    `json:"wfh"`
	IsActive bool      `json:"is_active"`
	Password string    `json:"password"`
	Freedays []string  `json:"free_days"`
}
