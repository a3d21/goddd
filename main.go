package main

import (
	"log"

	"github.com/a3d21/goddd/application"
	"github.com/a3d21/goddd/di"
)

func main() {
	c := di.GetContainerByEnv()

	err := c.Invoke(func(a *application.HttpApp) error {
		return a.Run()
	})

	if err != nil {
		log.Fatalf("fail to serve, err: %v", err)
	}
}
