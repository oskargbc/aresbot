package einhalb

type SecertDecode struct {
	Timestamp int    `json:"timestamp"`
	Action    string `json:"action"`
}

type CreateTokenResponse struct {
	PaymentID   string `json:"paymentID"`
	Success     bool   `json:"success"`
	RedirectURL string `json:"redirectUrl"`
}

type PaypalTokenResponse struct {
	Ack  string `json:"ack"`
	Data struct {
		Type  string `json:"type"`
		Token string `json:"token"`
	} `json:"data"`
	Meta struct {
		Calc string `json:"calc"`
		Rlog string `json:"rlog"`
	} `json:"meta"`
	Server string `json:"server"`
}
