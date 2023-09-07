package users

import (
	_ "embed"

	_ "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"eql/internal/app/entities/schema"
	"eql/internal/handlers"
)

type Endpoint interface {
	GetUserByID(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
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
// @Summary แสดงข้อมูล User ทั้งหมด
// @Description แสดงข้อมูลทั้งหมดแบบไม่มีเงื่อนไข
// @Accept  json
// @Produce  json
// @Success 200 {object} schema.HTTPError
// @Failure 400 {object} schema.HTTPError
// @Failure 404 {object} schema.HTTPError
// @Failure 500 {object} schema.HTTPError
// @Router /api/users/users2 [get]
// @Security ApiKeyAuth
func (ep *endpoint) GetUserByID(c *fiber.Ctx) error {

	return handlers.ResponseObjectNoRequest(c, ep.service.GetUserAll)
}

// @Tags User
// @Summary ค้นหา User ตามเงื่อนไข
// @Description Show User ตามเงื่อนไข
// @Accept  json
// @Produce  json
// @Param request query schema.UserRequest false " request body "
// @Success 200 {object} []models.User
// @Failure 400 {object} schema.HTTPError
// @Failure 404 {object} schema.HTTPError
// @Failure 500 {object} schema.HTTPError
// @Router /api/users/getuser [get]
// @Security ApiKeyAuth
func (ep *endpoint) GetUser(c *fiber.Ctx) error {
	return handlers.RespJson(c, ep.service.GetUser, &schema.UserRequest{})
}

// @Tags AUTH
// @Summary Login เข้าสู่ระบบ
// @Description login เข้าสู่ระบบเพื่อสร้าง jwt token
// @Accept  json
// @Produce  json
// @Param request body schema.LoginReq false " request body "
// @Success 200 {object} []models.User
// @Failure 400 {object} schema.HTTPError
// @Failure 404 {object} schema.HTTPError
// @Failure 500 {object} schema.HTTPError
// @Router /api/auth/login [post]
func (ep *endpoint) Login(c *fiber.Ctx) error {
	return handlers.RespJson(c, ep.service.GetUserByUserName, &schema.LoginReq{})
}
