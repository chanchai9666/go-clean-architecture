package main

import (
	_ "embed"
	"fmt"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"

	"eql/configs"
	"eql/docs"
	"eql/internal/handlers/router"
	"eql/internal/infrastructures/gofiber"
)

//go:embed files/swagger.json
var swaggerJSON []byte

func main() {
	conf, err := configs.GetConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
	app := gofiber.NewServer() //start go fiber
	// กำหนดค่า CORS ดังนี้
	app.Use(cors.New())
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	docs.SwaggerInfo.Host = "localhost:8318"
	docs.SwaggerInfo.Title = "EQL-SYSTEM"
	docs.SwaggerInfo.Description = "Test"
	docs.SwaggerInfo.Version = "0.1"
	// กำหนดรายละเอียดของส่วน auth Bearer
	// @securityDefinitions.apikey ApiKeyAuth
	// @name Authorization
	// @in ใส่ค่า Bearer เว้นวรรคและตามด้วย TOKEN  ex(Bearer ?????????)
	// END กำหนดค่าให้ swagger
	// =======================================================
	// เพิ่ม middleware สำหรับการเข้าถึง Swagger UI
	// เพิ่ม middleware สำหรับการเข้าถึง Swagger UI ด้วยควบคุมสิทธิ์
	app.Get("/swagger/*", swagger.New(swagger.Config{}))
	app.Get("/accounts/:id", ShowAccount)
	router.SetupRouter(app, c) //เรียก router ต่างๆ

	err = app.Listen(fmt.Sprintf(":%v", conf.AppConfig.Port))
	if err != nil {
		log.Fatalf("Error running server T^T: %s", err.Error())
	}

}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} Account
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /accounts/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	return c.JSON(Account{
		Id: c.Params("id"),
	})
}

type Account struct {
	Id string
}

type HTTPError struct {
	status  string
	message string
}
