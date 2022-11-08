package handlers

import "github.com/gofiber/fiber/v2"

type (
	Deposit interface {
		PostDepositMoney(ctx *fiber.Ctx) error
	}
)
