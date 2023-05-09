package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Audio(path string, router *fiber.App) {
	group :=router.Group(path)

	group.Get("/", func(c *fiber.Ctx) error { return c.SendString("OMG!!1")})

	group.Post("/", audioHandler)
}

func audioHandler(c *fiber.Ctx) error {
	fileName := c.Params("file")

	return c.SendFile(fileName)
}