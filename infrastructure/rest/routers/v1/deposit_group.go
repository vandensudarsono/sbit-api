package v1

import (
	"sbit/infrastructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

// InitDepositGroup: initialize group of deposit
func InitDepositGroup(router fiber.Router, handler handlers.Deposit) {
	depGroup := router.Group("/deposit")
	depGroup.Post("/post-deposit", handler.PostDepositMoney)
}
