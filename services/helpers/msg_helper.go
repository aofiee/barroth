package helpers

import (
	fiber "github.com/gofiber/fiber/v2"
)

func FailOnError(c *fiber.Ctx, err error, msg string, status int) error {
	if err != nil {
		c.Status(status).JSON(fiber.Map{
			"msg":   msg,
			"error": err.Error(),
		})
	}
	return nil
}
