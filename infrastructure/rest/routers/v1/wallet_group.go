package v1

import (
	"sbit/infrastructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

// InitWalletGroup: initialize group of wallet
func InitWalletGroup(router fiber.Router, handler handlers.Wallet) {
	walGroup := router.Group("/wallet")
	walGroup.Post("details", handler.GetBalance)
}
