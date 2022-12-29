package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

var service string

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

func init() {
	service = os.Getenv("SERVICE")
	if service == "" {
		log.Fatalln("You MUST set SERVICE env variable!")
	}
}

func main() {
	app := fiber.New()
	app.Get("/api/devices", getDevices)
	app.Listen(":8080")
}

func getDevices(c *fiber.Ctx) error {
	dvs := []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}

	return c.JSON(dvs)
}
