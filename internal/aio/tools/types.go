package tools

import "aresbot/internal/aio/shared"

type Response struct {
	e             error
	success       bool
	message       string
	publicWebhook shared.Webhook
	privatWebhook shared.Webhook
	logWebhook    shared.Webhook
}
