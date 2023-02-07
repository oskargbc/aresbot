package voostore

type CreatePaypalResponse struct {
	PaymentId        string `json:"paymentId"`
	HttpCacheEnabled bool   `json:"httpCacheEnabled"`
	CookieGroups     []struct {
		Name        string `json:"name"`
		Label       string `json:"label"`
		Description string `json:"description"`
		Cookies     []struct {
			Name            string `json:"name"`
			Label           string `json:"label"`
			GroupName       string `json:"groupName"`
			MatchingPattern string `json:"matchingPattern"`
		} `json:"cookies"`
		Required bool `json:"required"`
	} `json:"cookieGroups"`
	BilobaDetailSelectConfig struct {
		EnabledForPicture   bool `json:"enabled_for_picture"`
		EnabledForSelection bool `json:"enabled_for_selection"`
		EnabledForStandard  bool `json:"enabled_for_standard"`
		IgnoreNoStock       bool `json:"ignore_no_stock"`
		PreselectionQuery   bool `json:"preselectionQuery"`
		RedirectCode        int  `json:"redirect_code"`
	} `json:"bilobaDetailSelectConfig"`
}
