package public

import (
	"github.com/gofiber/fiber/v2"
)

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Response(code int, data interface{}, msg string, c *fiber.Ctx) error {
	err := c.JSON(Resp{
		code,
		data,
		msg,
	})
	return err
}
