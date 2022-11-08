package handlers

import "github.com/gofiber/fiber/v2"

type (
	Wallet interface {
		GetBalance(ctx *fiber.Ctx) error
	}
)
