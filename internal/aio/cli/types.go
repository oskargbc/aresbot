package cli

type HyperResponse struct {
	Account       string      `json:"account"`
	Email         string      `json:"email"`
	Unlocked      bool        `json:"unlocked"`
	Status        string      `json:"status"`
	CancelAt      interface{} `json:"cancel_at"`
	CanceledAt    interface{} `json:"canceled_at"`
	TransferPrice interface{} `json:"transfer_price"`
	Subscription  interface{} `json:"subscription"`
	Delinquent    bool        `json:"delinquent"`
	Plan          struct {
		Integrations struct {
			Discord struct {
				Guild        string   `json:"guild"`
				Roles        []string `json:"roles"`
				CancelAction string   `json:"cancel_action"`
			} `json:"discord"`
			Telegram interface{} `json:"telegram"`
		} `json:"integrations"`
		Transfers struct {
			Enabled      bool `json:"enabled"`
			CooldownDays int  `json:"cooldown_days"`
		} `json:"transfers"`
		Account          string        `json:"account"`
		Active           bool          `json:"active"`
		Name             string        `json:"name"`
		Image            string        `json:"image"`
		Description      string        `json:"description"`
		Amount           int           `json:"amount"`
		RentalPeriodDays interface{}   `json:"rental_period_days"`
		Currency         string        `json:"currency"`
		Recurring        interface{}   `json:"recurring"`
		Type             string        `json:"type"`
		Nft              interface{}   `json:"nft"`
		Files            []interface{} `json:"files"`
		Created          int64         `json:"created"`
		Links            []interface{} `json:"links"`
		Price            string        `json:"price"`
		Product          string        `json:"product"`
		Roles            []string      `json:"roles"`
		ID               string        `json:"id"`
	} `json:"plan"`
	Release   interface{} `json:"release"`
	Raffle    interface{} `json:"raffle"`
	Affiliate interface{} `json:"affiliate"`
	Coupon    interface{} `json:"coupon"`
	User      interface{} `json:"user"`
	Created   int64       `json:"created"`
	Key       string      `json:"key"`
	ID        string      `json:"id"`
	Metadata  struct {
	} `json:"metadata"`
	TrialEnd interface{} `json:"trial_end"`
}
