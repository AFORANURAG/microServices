package middleware

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

var Validate = validator.New()

type RequestBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type RequestParams struct {
	ID  int64 `json:"id" validate:"required"`
	AGE int   `json:"age"` // Added a missing tag
}

type RequestContainer struct {
	RequestBody   *RequestBody
	RequestParams *RequestParams
	RequestQuery  interface{}
}

func ValidateBody(field RequestContainer) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		customValidator := XValidator{validator: Validate} // Use the shared validator

		if field.RequestBody != nil {
			err := c.BodyParser(field.RequestBody)
			if err != nil {
				panic("error parsing body")
			}
			fmt.Printf("<---------------------------------The request body is-------------------------------------> %+v\n", field.RequestBody)

			if errs := customValidator.Validate(field.RequestBody); len(errs) > 0 && errs[0].Error {
				errMsgs := make([]string, len(errs))

				for i, err := range errs {
					errMsgs[i] = fmt.Sprintf(
						"[%s]: '%v' | Needs to implement '%s'",
						err.FailedField,
						err.Value,
						err.Tag,
					)
				}
				return &fiber.Error{
					Code:    fiber.ErrBadRequest.Code,
					Message: strings.Join(errMsgs, " and "),
				}
			}
		}

		if field.RequestParams != nil {
			err1 := c.ParamsParser(field.RequestParams)
			if err1 != nil {
				panic("error parsing params")
			}
			fmt.Printf("<---------------------------------The request Params are-------------------------------------> %+v\n", field.RequestParams)

			if errs := customValidator.Validate(field.RequestParams); len(errs) > 0 && errs[0].Error {
				errMsgs := make([]string, len(errs))

				for i, err := range errs {
					errMsgs[i] = fmt.Sprintf(
						"[%s]: '%v' | Needs to implement '%s'",
						err.FailedField,
						err.Value,
						err.Tag,
					)
				}
				return &fiber.Error{
					Code:    fiber.ErrBadRequest.Code,
					Message: strings.Join(errMsgs, " and "),
				}
			}
		}

		if field.RequestQuery != nil {
			err2 := c.QueryParser(field.RequestQuery)
			if err2 != nil {
				panic("error parsing query")
			}
			fmt.Printf("<---------------------------------The request query are-------------------------------------> %+v\n", field.RequestQuery)

			if errs := customValidator.Validate(field.RequestQuery); len(errs) > 0 && errs[0].Error {
				errMsgs := make([]string, len(errs))

				for i, err := range errs {
					errMsgs[i] = fmt.Sprintf(
						"[%s]: '%v' | Needs to implement '%s'",
						err.FailedField,
						err.Value,
						err.Tag,
					)
				}
				return &fiber.Error{
					Code:    fiber.ErrBadRequest.Code,
					Message: strings.Join(errMsgs, " and "),
				}
			}
		}

		return c.Next()
	}
}
