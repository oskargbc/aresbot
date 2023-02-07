package hyper

type bpStruct struct {
	UserAgent         string `json:"userAgent"`
	Timestamp         int64  `json:"timestamp"`
	IP                string `json:"ip"`
	Token             string `json:"token"`
	Auth              string `json:"auth"`
	Challenge         string `json:"challenge"`
	ChallengeFunction string `json:"challengeFunction"`
	CkHchKctl         bool   `json:"ck_hch_kctl"`
	CfCtl             string `json:"cf_ctl"`
	Domain            string `json:"domain"`
}
