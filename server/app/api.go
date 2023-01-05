package app

import (
	"pt-auto/service"

	"github.com/gofiber/fiber/v2"
)

func SetWebsiteConfig() fiber.Handler {
	return service.SetWebsiteConfig
}
