package helpers

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMsg(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := errors.New("test error")
		err := FailOnError(c, msg, "test error", 400)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		return nil
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	_, _ = app.Test(req)
}
