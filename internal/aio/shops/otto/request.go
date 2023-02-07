package otto

type AtcRequestBody struct {
	Services      []interface{} `json:"services"`
	Tracking      AtcTracking   `json:"tracking"`
	Articlenumber string        `json:"articleNumber"`
	Itemorigin    string        `json:"itemOrigin"`
	Quantity      string        `json:"quantity"`
	Variationid   string        `json:"variationId"`
}

type AtcTracking struct {
	ProductAvailabilityavg          interface{} `json:"product_AvailabilityAvg"`
	ProductAvailabilitymax          interface{} `json:"product_AvailabilityMax"`
	ProductAvailabilitymin          interface{} `json:"product_AvailabilityMin"`
	ProductAvailabilityregionaldiff interface{} `json:"product_AvailabilityRegionalDiff"`
	ProductAvailabilityview         string      `json:"product_AvailabilityView"`
}
