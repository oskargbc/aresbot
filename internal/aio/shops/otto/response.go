package otto

type AtcResponse struct {
	Addedquantity     int    `json:"addedQuantity"`
	Articlenumber     string `json:"articleNumber"`
	Itemcategory      string `json:"itemCategory"`
	Itemid            string `json:"itemId"`
	Itemname          string `json:"itemName"`
	Limitreached      bool   `json:"limitReached"`
	Requestedquantity int    `json:"requestedQuantity"`
	Uptolimit         bool   `json:"upToLimit"`
	Variationid       string `json:"variationId"`
}
