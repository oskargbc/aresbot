package models

type TaskGroup struct {
	Name  string
	Tasks []Task
}

type Task struct {
	ProfileName string `csv:"PROFILE"`
	ProductURL  string `csv:"URL"`
	Size        string `csv:"SIZE"`
	UseProxy    bool   `csv:"PROXY"`
	Mode        string `csv:"MODE"`
	Aco         bool   `csv:"ACO"`
	Region      string `csv:"REGION"`
	Store       string `csv:"STORE"`
	Keywords    string `csv:"KEYWORDS"`
	Sku         string `csv:"SKU"`
	Payment     string `csv:"PAYMENT"`
	Profile     Profile
	Id          int
}
