package middlewares

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type XValidator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type RequestComponent string

const (
	QueryParams RequestComponent = "queryParams"
	Body        RequestComponent = "body"
	Params      RequestComponent = "params"
)

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	errs := v.validator.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedField = err.StructField() // Use StructField() instead of Field()
			elem.Tag = err.Tag()                 // Export struct tag
			elem.Value = err.Value()             // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}

var validate = validator.New()

// rather than building a single function for universal validation, we can create request componentwise valdiations.
func ValidateRequest(requestComponent RequestComponent, schema interface{}) fiber.Handler {

	return func(c *fiber.Ctx) error {
		// Create a custom validator
		customValidator := XValidator{validator: validate}

		switch requestComponent {
		case "queryParams":
			if err := c.QueryParser(&schema); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
		case "body":
			if err := c.BodyParser(&schema); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
		case "params":
			if err := c.ParamsParser(&schema); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
		default:
			// Handle unknown requestComponent
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unknown RequestComponent"})
		}
		// Validate the schema using the custom validator
		fmt.Printf("schema: %+v\n", schema)
		if errs := customValidator.Validate(schema); len(errs) > 0 && errs[0].Error {
			errMsgs := make([]string, 0)
			for _, err := range errs {
				errMsgs = append(errMsgs, fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", err.FailedField, err.Value, err.Tag))
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": strings.Join(errMsgs, " and ")})
		}
		return c.Next()
	}
}
