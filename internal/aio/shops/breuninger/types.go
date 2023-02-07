package breuninger

type BreunigerArtikel struct {
	Initial bool `json:"initial"`
	Farben  struct {
		Aktiv struct {
			Name   string `json:"name"`
			FarbId string `json:"farbId"`
		} `json:"aktiv"`
		Werte []struct {
			Vorausgewaehlt bool   `json:"vorausgewaehlt"`
			ArtikelNr      string `json:"artikelNr"`
			Uri            string `json:"uri"`
			Wert           string `json:"wert"`
			Name           string `json:"name"`
			Outofstock     bool   `json:"outofstock"`
		} `json:"werte"`
	} `json:"farben"`
	Groessen struct {
		Aktiv struct {
			Name string `json:"name"`
		} `json:"aktiv"`
		Werte []struct {
			Name                   string `json:"name"`
			ArtikelNr              string `json:"artikelNr"`
			StockLowMessage        string `json:"stockLowMessage,omitempty"`
			StockLowMessageVariant string `json:"stockLowMessageVariant,omitempty"`
		} `json:"werte"`
	} `json:"groessen"`
	Preis struct {
		Schwarzpreis       string `json:"schwarzpreis"`
		SchemaPreis        string `json:"schemaPreis"`
		SchemaWaehrung     string `json:"schemaWaehrung"`
		MwstInfo           string `json:"mwstInfo"`
		SchemaAvailability string `json:"schemaAvailability"`
	} `json:"preis"`
	Ansichten []struct {
		StandardUrl               string `json:"standardUrl"`
		StandardRetinaUrl         string `json:"standardRetinaUrl"`
		StandardRedesignUrl       string `json:"standardRedesignUrl"`
		StandardRedesignRetinaUrl string `json:"standardRedesignRetinaUrl"`
		ZoomUrl                   string `json:"zoomUrl"`
		ZoomRetinaUrl             string `json:"zoomRetinaUrl"`
		ThumbnailUrl              string `json:"thumbnailUrl"`
		ThumbnailRetinaUrl        string `json:"thumbnailRetinaUrl"`
	} `json:"ansichten"`
	ArtikelNummer    string        `json:"artikelNummer"`
	VerkaufId        string        `json:"verkaufId"`
	Outofstock       bool          `json:"outofstock"`
	HvaUri           string        `json:"hvaUri"`
	Outfits          []interface{} `json:"outfits"`
	IsPartnerArtikel bool          `json:"isPartnerArtikel"`
}

type BreunigerAtcResponse struct {
	CartQuantityValue int `json:"cartQuantityValue"`
	Minicart          struct {
		Quantity             string `json:"quantity"`
		Total                string `json:"total"`
		CookieTotalName      string `json:"cookie_total_name"`
		CookieQuantityName   string `json:"cookie_quantity_name"`
		CookieTotalExpire    int    `json:"cookie_total_expire"`
		CookieQuantityExpire int    `json:"cookie_quantity_expire"`
	} `json:"minicart"`
	TrackingJson      []string `json:"trackingJson"`
	CubeEventTracking []struct {
		Currency struct {
			Code string `json:"code"`
		} `json:"currency"`
		EcomData struct {
			OrderValueLocal       int    `json:"order_value_local"`
			TaxValueLocal         int    `json:"tax_value_local"`
			ShippingValueLocal    int    `json:"shipping_value_local"`
			GiftWrappingCostLocal int    `json:"gift_wrapping_cost_local"`
			ListName              string `json:"list_name"`
		} `json:"ecom_data"`
		Items []struct {
			ProduktId       string `json:"produkt_id"`
			FarbId          string `json:"farb_id"`
			ArtikelId       string `json:"artikel_id"`
			VertriebsinfoId string `json:"vertriebsinfo_id"`
			Syan            string `json:"syan"`
			Quantity        int    `json:"quantity"`
		} `json:"items"`
		EventName               string `json:"event_name"`
		ResponsibleTeamTracking string `json:"responsible_team_tracking"`
	} `json:"cubeEventTracking"`
	Nachricht string `json:"nachricht"`
}
