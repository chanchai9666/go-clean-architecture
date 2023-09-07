package main

import (
	_ "embed"
	"fmt"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/valyala/fasthttp"

	"eql/configs"
	"eql/docs"
	"eql/internal/app/service"
	"eql/internal/handlers/router"
	"eql/internal/infrastructures/gofiber"
)

var swaggerJSON []byte

func main() {
	conf, err := configs.GetConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
	app := gofiber.NewServer() //start go fiber

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
	router.SetupRouter(app, c) //เรียก router ต่างๆ

	has, _ := service.HashPassword("12345")
	if service.CheckPassword("123456", has) {
		fmt.Println("Password OK")
	}

	err = app.Listen(fmt.Sprintf(":%v", conf.AppConfig.Port))
	if err != nil {
		log.Fatalf("Error running server T^T: %s", err.Error())
	}

}
