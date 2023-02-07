package mytoys

type AtcRequestBody struct {
	User       string `json:"user"`
	Auth       string `json:"auth"`
	Session    string `json:"session"`
	Client     string `json:"client"`
	Articlesku string `json:"articleSku"`
}
