package shops

import (
	"aresbot/internal/aio/models"
)

type Response struct {
	e           error
	rotateProxy bool
	success     bool
	message     string
	product     *models.Product
}
