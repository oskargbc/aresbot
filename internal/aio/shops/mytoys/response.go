package mytoys

type MytoysProduct struct {
	ArticleSku      interface{} `json:"articleSku"`
	ArticleGroupSku string      `json:"articleGroupSku"`
	BrandLogo       struct {
		Name     string `json:"name"`
		NameLink string `json:"nameLink"`
		Logo     struct {
			ImageURL string `json:"imageUrl"`
			Link     string `json:"link"`
		} `json:"logo"`
		URL interface{} `json:"url"`
	} `json:"brandLogo"`
	Rating struct {
		Stars   int `json:"stars"`
		Reviews int `json:"reviews"`
	} `json:"rating"`
	Images   []string      `json:"images"`
	Flags    []interface{} `json:"flags"`
	Discount float64       `json:"discount"`
	Title    string        `json:"title"`
	Authors  string        `json:"authors"`
	Price    struct {
		Default             string      `json:"default"`
		Reduced             interface{} `json:"reduced"`
		ReductionCode       interface{} `json:"reductionCode"`
		Old                 interface{} `json:"old"`
		IsBoundedStorePrice bool        `json:"isBoundedStorePrice"`
		HasMinPrice         bool        `json:"hasMinPrice"`
		AdditionalInfo      interface{} `json:"additionalInfo"`
		Currency            string      `json:"currency"`
	} `json:"price"`
	SalesCategoryUpper string        `json:"salesCategoryUpper"`
	SalesCategoryLower string        `json:"salesCategoryLower"`
	EnergyImages       []interface{} `json:"energyImages"`
	OnlyShipTo         interface{}   `json:"onlyShipTo"`
	Availability       struct {
		Code            string      `json:"code"`
		PublicationDate string      `json:"publicationDate"`
		MinDeliveryDays int         `json:"minDeliveryDays"`
		MaxDeliveryDays int         `json:"maxDeliveryDays"`
		AvailableItems  interface{} `json:"availableItems"`
	} `json:"availability"`
	PaybackPoints int `json:"paybackPoints"`
	Description   struct {
		Text        string        `json:"text"`
		Awards      []interface{} `json:"awards"`
		Warnings    interface{}   `json:"warnings"`
		Publication struct {
			Date      interface{} `json:"date"`
			Publisher interface{} `json:"publisher"`
			Binding   interface{} `json:"binding"`
		} `json:"publication"`
		Isbn     interface{} `json:"isbn"`
		RunTime  interface{} `json:"runTime"`
		GameInfo struct {
			PlayTime        interface{} `json:"playTime"`
			NumOfPieces     interface{} `json:"numOfPieces"`
			MinNumOfPlayers interface{} `json:"minNumOfPlayers"`
			MaxNumOfPlayers interface{} `json:"maxNumOfPlayers"`
		} `json:"gameInfo"`
		Age struct {
			MinMonth interface{} `json:"minMonth"`
			MaxMonth interface{} `json:"maxMonth"`
			MinYear  string      `json:"minYear"`
			MaxYear  string      `json:"maxYear"`
		} `json:"age"`
		ArticleNumber string `json:"articleNumber"`
	} `json:"description"`
	Publication               interface{}   `json:"publication"`
	Size                      interface{}   `json:"size"`
	Sizes                     []interface{} `json:"sizes"`
	Color                     interface{}   `json:"color"`
	Colors                    []interface{} `json:"colors"`
	ActionCodes               interface{}   `json:"actionCodes"`
	MediaFiles                interface{}   `json:"mediaFiles"`
	SupplierName              interface{}   `json:"supplierName"`
	RecommendedProductsTitle  interface{}   `json:"recommendedProductsTitle"`
	RecommendedProducts       interface{}   `json:"recommendedProducts"`
	OtherClientsProductsTitle interface{}   `json:"otherClientsProductsTitle"`
	OtherClientsProducts      interface{}   `json:"otherClientsProducts"`
	ProductBundles            []interface{} `json:"productBundles"`
	DeliveryMessage           interface{}   `json:"deliveryMessage"`
	QuestionsAnswersID        interface{}   `json:"questionsAnswersId"`
}

type SessionResponse struct {
	Msuser              string `json:"msUser"`
	Msauth              string `json:"msAuth"`
	Mssession           string `json:"msSession"`
	Msclient            string `json:"msClient"`
	Msuserexpiration    int    `json:"msUserExpiration"`
	Msauthexpiration    int    `json:"msAuthExpiration"`
	Mssessionexpiration int    `json:"msSessionExpiration"`
	Msclientexpiration  int    `json:"msClientExpiration"`
	Empty               bool   `json:"empty"`
}

type AtcResponse struct {
	Articlecount  int               `json:"articleCount"`
	Resulttype    string            `json:"resultType"`
	Lineitemmodel map[string]string `json:"results"`
}

type BasketResponse struct {
	Basket struct {
		Goodscounttotal       int  `json:"goodsCountTotal"`
		Removeditemscount     int  `json:"removedItemsCount"`
		Containsmytoyscatalog bool `json:"containsMyToysCatalog"`
		Environmentpopup      bool `json:"environmentPopup"`
		Deliverygroups        struct {
			Num0 struct {
				Headerinfo    string           `json:"headerInfo"`
				Shoppingitems []MyToysCartItem `json:"shoppingItems"`
			} `json:"0"`
		} `json:"deliveryGroups"`
		Baskettotals struct {
			Ordervalue struct {
				Price  string `json:"price"`
				Gratis bool   `json:"gratis"`
			} `json:"orderValue"`
			Costsandrebates struct {
			} `json:"costsAndRebates"`
			Total struct {
				Price  string `json:"price"`
				Gratis bool   `json:"gratis"`
			} `json:"total"`
			Multishopfreeshipping bool `json:"multishopFreeShipping"`
		} `json:"basketTotals"`
		Messages struct {
		} `json:"messages"`
	} `json:"basket"`
	Trackingdata struct {
		Tcvars string `json:"tcVars"`
	} `json:"trackingData"`
}

type MyToysCartItem struct {
	Productimageurl    string `json:"productImageUrl"`
	Pdpurl             string `json:"pdpUrl"`
	Quantity           int    `json:"quantity"`
	Singleproductprice struct {
		Price  string `json:"price"`
		Gratis bool   `json:"gratis"`
	} `json:"singleProductPrice"`
	Singleprices struct {
		Price  string `json:"price"`
		Gratis bool   `json:"gratis"`
	} `json:"singlePrices"`
	Hasvariants   bool   `json:"hasVariants"`
	Commission    bool   `json:"commission"`
	Removed       bool   `json:"removed"`
	Fixed         bool   `json:"fixed"`
	Innotepad     bool   `json:"inNotepad"`
	Color         string `json:"color"`
	Brand         string `json:"brand"`
	Catalogdomain string `json:"catalogDomain"`
	Sku           string `json:"sku"`
	Description   string `json:"description"`
	Position      string `json:"position"`
	Messages      struct {
	} `json:"messages"`
	Shopimage string `json:"shopImage"`
}

type BasketATCResponse struct {
	User struct {
		Authenticated bool   `json:"authenticated"`
		Firstname     string `json:"firstName"`
		Lastname      string `json:"lastName"`
	} `json:"user"`
	Basket struct {
		Articlecount int `json:"articleCount"`
	} `json:"basket"`
	Notepad struct {
		Articlecount int `json:"articleCount"`
	} `json:"notepad"`
}

type ItemDeleteResponse struct {
	Resultstatus int `json:"resultStatus"`
}

type ApiOrderResponse struct {
	Forwardto string `json:"forwardTo"`
}
