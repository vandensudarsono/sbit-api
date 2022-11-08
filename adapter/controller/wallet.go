package controller

import "github.com/gofiber/fiber/v2"

type (
	WalletController struct{}

	BalanceRequest struct {
		WalletID int `json:"wallet_id"`
	}

	BalanceResponse struct {
		WalletID      int     `json:"wallet_id"`
		Balance       float64 `json:"balance"`
		AboveTreshold bool    `json:"above_treshold"`
	}
)

// NewWallterController: create wallet controller
func NewWalletController() *WalletController {
	return &WalletController{}
}

// GetBalnce: get balance controller
func (*WalletController) GetBalance(ctx *fiber.Ctx) error {
	var (
		balanceRequest  BalanceRequest
		balanceResponse BalanceResponse
		err             error
	)

	err = ctx.BodyParser(&balanceRequest)
	if err != nil {
		//log error
		return err
	}

	//validate

	//send to usecase

	return ctx.JSON(balanceResponse)
}
