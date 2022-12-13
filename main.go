package main

import (
	"Monitoreo/Utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

func main() {
	Utils.QueryAuth.ClientId = os.Getenv("APP-ID")
	Utils.QueryAuth.ClientSecret = os.Getenv("APP-Secret")
	Utils.QueryAuth.Scope = os.Getenv("APP-Scope")
	Utils.UipathOrg.Id = os.Getenv("Orchestrator-ID")

	Utils.LastMonitoredTime = time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC)
	_ = Utils.RefreshUiPathToken()
	go func() {
		for {
			err := Utils.RefreshUiPathToken()
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(59 * time.Minute) // Refresh the token every 10 minutes
		}
	}()
	ProcessResponse := new(Utils.ProcessResponse)
	err := ProcessResponse.Get(3902201)
	if err != nil {
		fmt.Println(err)
	}
	Processes := ProcessResponse.GetProcesses()
	LogsResponse := new(Utils.LogResponse)
	err = LogsResponse.Get(3902201)
	if err != nil {
		fmt.Println(err)
	}
	Logs := LogsResponse.FilterLogs(Processes)
	fmt.Println(Logs)
	app := fiber.New()
	app.Get("/AuthReturn", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
