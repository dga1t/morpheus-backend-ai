package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"

	"morpheus-backend-ai/src/controllers"
)

// TODO
// - receive audio from Unity server (discuss how - http/ws ??) 
// 			* If possible - use a lossless encoding such as FLAC or LINEAR16
// 			* If possible - capture audio using a sample rate of 16000 Hz (or higher)

// - send request to GCP speech-to-text api
// - make requests to GPT api
// - GCP text-to-speech

func main() {
	fmt.Println(`OMG!!1`)

	app := fiber.New(fiber.Config{
		BodyLimit:    200 * 1024 * 1024,
	})
	app.Use(cors.New())

	controllers.Audio("/audio", app)

	app.Listen(":3000")
}