package users

import (
	_ "embed"

	_ "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"eql/internal/handlers"
)

type Endpoint interface {
	GetUserByID(c *fiber.Ctx) error
}

type endpoint struct {
	service Service
}

func NewEndpoint(service Service) Endpoint {
	return &endpoint{
		service: service,
	}
}

// @Tags User
// @Summary Show User
// @Description Show All User
// @Accept  json
// @Produce  json
// @Success 200 {object} schema.HTTPError
// @Failure 400 {object} schema.HTTPError
// @Failure 404 {object} schema.HTTPError
// @Failure 500 {object} schema.HTTPError
// @Router /api/users/users2 [get]
func (ep *endpoint) GetUserByID(c *fiber.Ctx) error {
	return handlers.ResponseObjectNoRequest(c, ep.service.GetUser)
}
