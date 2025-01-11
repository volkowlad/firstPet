package main

import (
	"firstPet/internal/pkg/app"

	"github.com/sirupsen/logrus"
)

func main() {
	a, err := app.New()
	if err != nil {
		logrus.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
