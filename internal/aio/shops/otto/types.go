package otto

type OttoProduct struct {
	Description            interface{}              `json:"description"`
	Distinctdimensions     []interface{}            `json:"distinctDimensions"`
	Offerid                interface{}              `json:"offerId"`
	Sortedvariationids     []string                 `json:"sortedVariationIds"`
	Variations             map[string]OttoVariation `json:"variations"`
	Variationtree          interface{}              `json:"variationTree"`
	Brand                  string                   `json:"brand"`
	Brandimageid           string                   `json:"brandImageId"`
	Gratisinformationvalue string                   `json:"gratisInformationValue"`
	Hasoptions             bool                     `json:"hasOptions"`
	Htmlcharacteristics    string                   `json:"htmlCharacteristics"`
	ID                     string                   `json:"id"`
}

type OttoVariation struct {
	Availability VariationAvailability `json:"availability"`
	Benefits     []struct {
		ID                         string      `json:"id"`
		Formatteddiscountprice     string      `json:"formattedDiscountPrice"`
		Formattednormdiscountprice interface{} `json:"formattedNormDiscountPrice"`
		Discountvaluepercent       string      `json:"discountValuePercent"`
	} `json:"benefits"`
	Businesscategory struct {
		Group string `json:"group"`
		Value string `json:"value"`
	} `json:"businessCategory"`
	Cashback          interface{} `json:"cashback"`
	Consultingdetails []struct {
		Displayname string `json:"displayName"`
		Name        string `json:"name"`
		Contenturl  string `json:"contentUrl"`
	} `json:"consultingDetails"`
	Customdimensions struct {
		Dimension []interface{} `json:"dimension"`
	} `json:"customDimensions"`
	Custommeasuretype interface{} `json:"customMeasureType"`
	Deal              interface{} `json:"deal"`
	Deliveryinfo      struct {
		Shippingflags            []string `json:"shippingFlags"`
		Nextpossibledeliverydate string   `json:"nextPossibleDeliveryDate"`
	} `json:"deliveryInfo"`
	Description interface{} `json:"description"`
	Detailicons struct {
		Care    []interface{} `json:"care"`
		Cut     []interface{} `json:"cut"`
		Quality []interface{} `json:"quality"`
	} `json:"detailIcons"`
	Dimensions struct {
		Dimension []interface{} `json:"dimension"`
	} `json:"dimensions"`
	Displayprice struct {
		Comparativepriceadvantage interface{} `json:"comparativePriceAdvantage"`
		Comparativepriceamount    interface{} `json:"comparativePriceAmount"`
		Custommeasuretype         interface{} `json:"customMeasureType"`
		Installments              struct {
			Amount string `json:"amount"`
			Count  int    `json:"count"`
		} `json:"installments"`
		Normprice                          interface{} `json:"normPrice"`
		Pricedifferencetofirstvariation    interface{} `json:"priceDifferenceToFirstVariation"`
		Formattedpriceamount               string      `json:"formattedPriceAmount"`
		Hasmoreexpensiveavailablevariation bool        `json:"hasMoreExpensiveAvailableVariation"`
		Hassuggestedretailprice            bool        `json:"hasSuggestedRetailPrice"`
		Priceamount                        int         `json:"priceAmount"`
		Techpriceamount                    string      `json:"techPriceAmount"`
	} `json:"displayPrice"`
	Downloadabledocuments []interface{} `json:"downloadableDocuments"`
	Energyefficiency      interface{}   `json:"energyEfficiency"`
	Energyefficiencyclass interface{}   `json:"energyEfficiencyClass"`
	Expertreviews         interface{}   `json:"expertReviews"`
	Htmlcharacteristics   interface{}   `json:"htmlCharacteristics"`
	Images                []struct {
		ID              string `json:"id"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
		Mainimage       bool   `json:"mainImage"`
		Initializeimage bool   `json:"initializeImage"`
		Selected        bool   `json:"selected"`
		Number          int    `json:"number"`
		Index           int    `json:"index"`
	} `json:"images"`
	Links               interface{} `json:"links"`
	Manufacturercontent interface{} `json:"manufacturerContent"`
	Patternsample       interface{} `json:"patternSample"`
	Retailer            struct {
		ID     string `json:"id"`
		Isotto bool   `json:"isOtto"`
		Name   string `json:"name"`
	} `json:"retailer"`
	Richcontenturl interface{} `json:"richContentUrl"`
	Score          interface{} `json:"score"`
	Sellingpoints  struct {
		Sellingpoint []string `json:"sellingPoint"`
	} `json:"sellingPoints"`
	Skusuffix                  interface{} `json:"skuSuffix"`
	Sustainability             interface{} `json:"sustainability"`
	Articlenumber              string      `json:"articleNumber"`
	Articlenumberwithpromotion string      `json:"articleNumberWithPromotion"`
	Assortmentcode             string      `json:"assortmentCode"`
	Businessmodel              string      `json:"businessModel"`
	Companyid                  int         `json:"companyId"`
	Dimensioncount             int         `json:"dimensionCount"`
	Ean                        string      `json:"ean"`
	Hazmat                     bool        `json:"hazmat"`
	ID                         string      `json:"id"`
	Mainimagerelativeheight    int         `json:"mainImageRelativeHeight"`
	Moin                       string      `json:"moin"`
	Name                       string      `json:"name"`
	Productid                  string      `json:"productId"`
	Promotionnumber            string      `json:"promotionNumber"`
	Sale                       bool        `json:"sale"`
	Seoimageurlcurrent         string      `json:"seoImageUrlCurrent"`
	Sustainable                bool        `json:"sustainable"`
	Title                      string      `json:"title"`
}

type VariationAvailability struct {
	Avgdeliverytime             interface{} `json:"avgDeliveryTime"`
	Avgdeliverytimeregionaldiff interface{} `json:"avgDeliveryTimeRegionalDiff"`
	Limitation                  interface{} `json:"limitation"`
	Maxdeliverytime             interface{} `json:"maxDeliveryTime"`
	Mindeliverytime             interface{} `json:"minDeliveryTime"`
	Buyable                     bool        `json:"buyable"`
	Displayname                 string      `json:"displayName"`
	Limited                     bool        `json:"limited"`
	Orderable                   bool        `json:"orderable"`
	Showcoronadeliverydelayinfo bool        `json:"showCoronaDeliveryDelayInfo"`
	Status                      string      `json:"status"`
}
