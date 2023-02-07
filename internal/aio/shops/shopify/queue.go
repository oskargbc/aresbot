package shopify

import "time"

type QueuePollRequestBody struct {
	Query     string                        `json:"query"`
	Variables QueuePollRequestBodyVariables `json:"variables"`
}

type QueuePollRequestBodyVariables struct {
	Token string `json:"token"`
}

type QueueRollResponseBody struct {
	Data QueueRollResponseBodyData `json:"data"`
}

type QueueRollResponseBodyData struct {
	Poll QueueRollResponsePollData `json:"poll"`
}

type QueueRollResponsePollData struct {
	Token                      string                    `json:"token"`
	PollAfter                  time.Time                 `json:"pollAfter"`
	QueueEtaSeconds            int                       `json:"queueEtaSeconds"`
	ProductVariantAvailability []PollProductAvailability `json:"productVariantAvailability"`
	Typename                   string                    `json:"__typename"`
}

type PollProductAvailability struct {
	ID        int64 `json:"id"`
	Available bool  `json:"available"`
}
