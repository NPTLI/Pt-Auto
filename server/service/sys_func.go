package service

import (
	"pt-auto/public"

	"github.com/gofiber/fiber/v2"
)

func SetWebsiteConfig(c *fiber.Ctx) error {

	err := public.Response(200, map[string]string{}, "获取成功", c)
	return err
}
