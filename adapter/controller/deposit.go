package controller

import (
	"github.com/gofiber/fiber/v2"
)

type (
	DepositController struct {
	}

	DepositMoney struct {
		Wallet_id int     `json:"wallet_id"`
		Amount    float64 `json:"Amount"`
	}
)

// NewDepositController: Create deposit controllerß
func NewDepositController() *DepositController {
	return &DepositController{}
}

func (*DepositController) PostDepositMoney(ctx *fiber.Ctx) error {
	var (
		depositRequest DepositMoney
		err            error
	)

	err = ctx.BodyParser(&depositRequest)
	if err != nil {
		//log error
		return err
	}

	//validate

	//send to usecase

	return ctx.JSON(depositRequest)
}
