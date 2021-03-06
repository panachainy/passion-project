package covid

import (
	"fmt"
	"sync"

	"covid-19-api/internal/platform/storage/cache"

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
	Cache   cache.CacheInterface
}

func ProviderHandler(s CovidService, c cache.CacheInterface) *CovidHandlerImp {
	hOnce.Do(
		func() {
			handler = &CovidHandlerImp{
				Service: s,
				Cache:   c,
			}
		},
	)

	return handler
}

func (h *CovidHandlerImp) GetToday(c *fiber.Ctx) error {
	var covidSerializer CovidSerializer
	if err := h.Cache.Get(TODAY_CACHE_KEY, &covidSerializer); err == nil {
		return c.JSON(fiber.Map{
			"data":    covidSerializer,
			"message": "success",
		})
	}

	res, err := h.Service.GetToday()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable,
			fmt.Sprintf("GetToday 2: %v", err.Error()))
	}

	_ = h.Cache.Set(TODAY_CACHE_KEY, res[0], CACHE_TIME)

	return c.JSON(fiber.Map{
		"data":    res[0],
		"message": "success",
	})
}
