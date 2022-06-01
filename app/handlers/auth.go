package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Login service logs in a user
func Login(ctx *fiber.Ctx) error {

	//TODO:: Check user's email
	//			Verify Assword
	//			Update session
	//

	return nil
}

// Logout service logs out a user
func Logout(ctx *fiber.Ctx) error {
	//TODO:: Destroy session

	return nil
}
