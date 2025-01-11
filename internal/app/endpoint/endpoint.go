package endpoint

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(c *fiber.Ctx) error {
	days := e.s.DaysLeft()

	res := fmt.Sprintf("Количество дней до запуска приложения: %d", days)

	err := c.Status(fiber.StatusOK).SendString(res)
	if err != nil {
		return err
	}

	return nil
}
