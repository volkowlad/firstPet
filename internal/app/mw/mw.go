package mw

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RoleCheck(c *fiber.Ctx) error {
	val := string(c.Request().Header.Peek("User-Role"))

	if strings.ToLower(val) == "volkov" {
		logrus.Println("Разработчик")
	}

	return c.Next()
}
