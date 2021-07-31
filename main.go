package main

import (
	"e-comm-auth/handler"
	"e-comm-auth/model"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiberConf)
	app.Post("sign-up", handler.Signup)
	app.Listen(":8080")
}

var fiberConf fiber.Config = fiber.Config{
	AppName:      "E-Comm Auth",
	ErrorHandler: errorHandler,
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	e, ok := err.(*fiber.Error)
	if !ok {
		return unknownError(ctx)
	}

	resp := new(model.ErrorResponse)
	resp.CreateErrorResponse(e)
	err = ctx.Status(resp.Code).JSON(resp)
	if err != nil {
		return unknownError(ctx)
	}
	return nil
}

func unknownError(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
}
