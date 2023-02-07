package models

type Settings struct {
	MonitorDelay       int      `json:"monitorDelay"`
	RetryCount         int      `json:"retryCount"`
	WebhookURL         string   `json:"webhookUrl"`
	License            string   `json:"license"`
	ProxyHttps         bool     `json:"proxyHttps"`
	CaptchaProvider    string   `json:"captchaProvider"`
	CaptchaKey         string   `json:"captchaKey"`
	QuicktaskProfile   string   `json:"quicktaskProfile"`
	QuicktaskAmount    int      `json:"quicktaskAmount"`
	QuicktaskWithProxy bool     `json:"quicktaskWithProxy"`
	DiscordToken       string   `json:"discordToken"`
	DiscordChannelIds  []string `json:"discordChannelIds"`
	DiscordServerIds   []string `json:"discordServerIds"`
	DiscordKeywords    []string `json:"discordKeywords"`
}
