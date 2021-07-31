package handler

import (
	"e-comm-auth/db"
	"e-comm-auth/model"
	"e-comm-auth/utils"
	"github.com/gofiber/fiber/v2"
)

type handlerFunc func(ctx *fiber.Ctx) error

var Signup handlerFunc = func(ctx *fiber.Ctx) error {
	userModel := new(model.UserSignupModel)
	if err := ctx.BodyParser(userModel); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid Request Body"}
	}

	/** Validating fields in the payload */
	if utils.IsEmpty(userModel.Name, userModel.LastName) || !utils.IsEmailValid(userModel.Email) || !utils.IsGreaterThanEq(userModel.Password, 6) {
		return &fiber.Error{Code: 400, Message: "Validation Failed!"}
	}

	/** Checking if user already exists */
	if db.UserExists(userModel.Email) {
		return &fiber.Error{Code: 400, Message: "User already signed-up"}
	}
	user := userModel.ToUser()

	/**
	 * Inserting user
	 * Register user to Firebase
	 */
	err := db.InsertUser(user, userModel)
	if err != nil {
		return &fiber.Error{Code: 500, Message: "User could not signed-up!"}
	}

	utils.SendStructResponse(ctx, user, 201)
	return nil
}
