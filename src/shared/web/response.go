package web

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type Response struct {
	Data   any             `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func OKResponse(ctx *fiber.Ctx, status int, data any) error {
	return ctx.Status(status).JSON(Response{
		Data: data,
	})
}

func ErrResponse(ctx *fiber.Ctx, err error) error {
	validateErr, ok := errors.Cause(err).(validator.ValidationErrors)
	if ok {
		var errs []ErrorResponse
		trans, _ := translator.GetTranslator("en")
		for _, erritem := range validateErr {

			errs = append(errs, ErrorResponse{
				Code:    400,
				Message: erritem.Translate(trans),
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(Response{
			Errors: errs,
		})
	}
	fiberErr, ok := errors.Cause(err).(*fiber.Error)
	if ok {
		return ctx.Status(fiberErr.Code).JSON(Response{
			Errors: []ErrorResponse{
				{
					Code:    fiberErr.Code,
					Message: fiberErr.Message,
				},
			},
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Errors: []ErrorResponse{
			{
				Code:    500,
				Message: err.Error(),
			},
		},
	})
}
