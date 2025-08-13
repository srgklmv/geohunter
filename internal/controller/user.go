package controller

import (
	"github.com/gofiber/fiber/v3"
)

/*
	The user's controller should include registration, login and account
	management logic with following handlers:
	1. registration
	2. activate account
	3. login
	4. logout
	5. refresh token

	6. get user
	7. update user
	8. delete user
*/

func (c *Controller) Registration(fc *fiber.Ctx) error {
	// TODO: Request body.

	// TODO: Input validation.

	// TODO: Validation for existing account.

	// TODO: Create account.

	// TODO: Create tokens.

	// TODO: Return.

	return nil
}

func (c *Controller) ActivateAccount(fc *fiber.Ctx) error {
	panic("not implemented")
}

func (c *Controller) Login(fc *fiber.Ctx) error {
	panic("not implemented")
}

func (c *Controller) Logout(fc *fiber.Ctx) error {
	panic("not implemented")
}

func (c *Controller) RefreshToken(fc *fiber.Ctx) error {
	panic("not implemented")
}

func (c *Controller) GetUser(fc *fiber.Ctx) error {
	panic("not implemented")
}

func (c *Controller) UpdateUser(fc *fiber.Ctx) error {
	panic("not implemented")
}

func (c *Controller) DeleteUser(fc *fiber.Ctx) error {
	panic("not implemented")
}
