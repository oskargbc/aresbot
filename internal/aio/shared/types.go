package shared

import (
	models2 "aresbot/internal/aio/models"
)

type Webhook struct {
	Content   string      `json:"content"`
	Username  string      `json:"username"`
	AvatarURL interface{} `json:"avatar_url"`
	Tts       bool        `json:"tts"`
	Embeds    []Embeds    `json:"embeds"`
}
type Footer struct {
	Text    string      `json:"text"`
	IconURL interface{} `json:"icon_url"`
}
type Image struct {
	URL interface{} `json:"url"`
}
type Thumbnail struct {
	URL string `json:"url"`
}
type Author struct {
	Name    string      `json:"name"`
	URL     interface{} `json:"url"`
	IconURL interface{} `json:"icon_url"`
}
type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type Embeds struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	URL         string      `json:"url"`
	Timestamp   interface{} `json:"timestamp"`
	Color       int         `json:"color"`
	Footer      Footer      `json:"footer"`
	Image       Image       `json:"image"`
	Thumbnail   Thumbnail   `json:"thumbnail"`
	Author      Author      `json:"author"`
	Fields      []Fields    `json:"fields"`
}

type WebhookData struct {
	Product models2.Product
	Profile models2.Profile
}

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

type ApiVersionResponse struct {
	Version string `json:"version"`
}

type ApiStatusResponse struct {
	Online bool   `json:"online"`
	Status string `json:"status"`
}

type CliData struct {
	Hyper          HyperResponse
	Version        string
	Groups         []models2.TaskGroup
	Settings       models2.Settings
	Proxys         []models2.Proxy
	WebhookHandler WebhookHandler
	Shops          []string
	Tools          []string
}

type CliResponse struct {
	Extras   map[string]interface{}
	ToolName string
	Tasks    []models2.Task
	Cmd      *string
}
