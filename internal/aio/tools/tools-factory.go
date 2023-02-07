package tools

import (
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
)

type Tool interface {
	Run(rotator shared.ProxyRotator, extras map[string]interface{}, handler shared.WebhookHandler)
	IsActive() bool
	GetSettings(settings models.Settings, id int) error
	NeedBackend() bool
}

var AllTools = map[string]Tool{
	"QueueIt": NewQueueItTool(),
	"Mobile":  NewMobileTool(),
}

/*
// new tool
func NewQueueItTool() Tool {
	return &QueueItTool{}
}

type QueueItTool struct {
	Tool
}

func (s *QueueItTool) Run() {

}

func (s *QueueItTool) IsActive() bool {
	return true
}

func (s *QueueItTool) NeedsProxys() bool {
	return false
}

func (s *QueueItTool) GetSettings() error {
	return nil
}

func (s *QueueItTool) NeedBackend() bool {
	return false
}

*/
