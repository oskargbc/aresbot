package tools

import (
	"aresbot/internal/aio/models"
	"aresbot/internal/aio/shared"
	mobile_socket "aresbot/internal/aio/tools/mobile-socket"
)

func NewMobileTool() *MobileTool {
	return &MobileTool{}
}

type MobileTool struct {
	Tool
}

func (s *MobileTool) Run(rotator shared.ProxyRotator, extras map[string]interface{}, handler shared.WebhookHandler) {
	m := mobile_socket.NewMobileSocket()
	m.Run()
	return

}

func (s *MobileTool) IsActive() bool {
	return true
}

func (s *MobileTool) GetSettings(settings models.Settings, id int) error {
	return nil
}

func (s *MobileTool) NeedBackend() bool {
	return false
}
