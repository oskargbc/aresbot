package models

type Profile struct {
	ProfileName string `csv:"profileName"`
	FirstName   string `csv:"firstName"`
	LastName    string `csv:"lastName"`
	Email       string `csv:"email"`
	Phone       string `csv:"phone"`
	Address     string `csv:"address"`
	Address2    string `csv:"address2"`
	Zip         string `csv:"zip"`
	City        string `csv:"city"`
	Country     string `csv:"country"`
	State       string `csv:"state"`
	Cardname    string `csv:"cardName"`
	Cnb         string `csv:"cnb"`
	Month       string `csv:"month"`
	Year        string `csv:"year"`
	Cvv         string `csv:"cvv"`
	CardType    string `csv:"cardType"`
	Password    string `csv:"password"`
}
