package covid

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	handler *CovidHandlerImp
	hOnce   sync.Once
)

type CovidHandler interface {
	GetToday(c *fiber.Ctx) error
}

type CovidHandlerImp struct {
	Service CovidService
}

func ProviderHandler(s CovidService) *CovidHandlerImp {
	hOnce.Do(
		func() {
			handler = &CovidHandlerImp{
				Service: s,
			}
		},
	)

	return handler
}

func (h *CovidHandlerImp) GetToday(c *fiber.Ctx) error {
	h.Service.GetToday()

	return c.SendString("up")
}
