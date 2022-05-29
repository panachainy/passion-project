package covid

import (
	"fmt"
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
	res, err := h.Service.GetToday()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable,
			fmt.Sprintf("GetToday 2: %v", err.Error()))
	}

	return c.JSON(fiber.Map{
		"data":    res[0],
		"message": "success",
	})
}
