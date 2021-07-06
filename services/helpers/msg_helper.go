package helpers

import (
	"github.com/go-playground/validator"
	fiber "github.com/gofiber/fiber/v2"
)

type (
	ErrorResponse struct {
		FailedField string
		Tag         string
		Value       string
	}
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
func ValidateStruct(v interface{}) []*ErrorResponse {
	var errs []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(v)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errs = append(errs, &element)
		}
	}
	return errs
}
