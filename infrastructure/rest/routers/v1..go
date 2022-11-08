package routers

import "github.com/gofiber/fiber/v2"

// InitV1: initalze v1 route endpoint
func InitV1(router fiber.Router) fiber.Router {
	var route string = "/v1"
	return router.Group(route)
}
