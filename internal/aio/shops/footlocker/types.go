package footlocker

type ShippingPayload struct {
	ShippingAddress struct {
		SetAsDefaultBilling  bool   `json:"setAsDefaultBilling"`
		SetAsDefaultShipping bool   `json:"setAsDefaultShipping"`
		FirstName            string `json:"firstName"`
		LastName             string `json:"lastName"`
		Email                bool   `json:"email"`
		Phone                string `json:"phone"`
		Country              struct {
			Isocode string `json:"isocode"`
			Name    string `json:"name"`
		} `json:"country"`
		ID                interface{} `json:"id"`
		SetAsBilling      bool        `json:"setAsBilling"`
		SaveInAddressBook bool        `json:"saveInAddressBook"`
		Type              string      `json:"type"`
		LoqateSearch      string      `json:"LoqateSearch"`
		ShippingAddress   bool        `json:"shippingAddress"`
	} `json:"shippingAddress"`
}

type Product struct {
	Name          string `json:"name"`
	IsSaleProduct bool   `json:"isSaleProduct"`
	Image         string
	Images        []Images        `json:"images"`
	SellableUnits []SellableUnits `json:"sellableUnits"`
}

type Images struct {
	Variations []Variant `json:"variations"`
}

type SellableUnits struct {
	Price         Price       `json:"price"`
	IsRecaptchaOn bool        `json:"isRecaptchaOn"`
	Attributes    []Attribute `json:"attributes"`
	Status        string      `json:"stockLevelStatus"`
}

type Price struct {
	OriginalPrice float64 `json:"originalPrice"`
	CurrentPrice  float64 `json:"value"`
}

type Attribute struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Variant struct {
	Formant string `json:"format"`
	Url     string `json:"url"`
}
