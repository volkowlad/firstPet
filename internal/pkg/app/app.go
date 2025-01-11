package app

import (
	"firstPet/internal/app/endpoint"
	"firstPet/internal/app/mw"
	"firstPet/internal/app/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type App struct {
	e      *endpoint.Endpoint
	s      *service.Service
	webApp *fiber.App
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.webApp = fiber.New()

	a.webApp.Use(mw.RoleCheck)

	a.webApp.Get("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("server running")

	logrus.Fatal(a.webApp.Listen(":8080"))

	return nil
}
