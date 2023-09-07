// handlers/router/router.go
package router

import (
	"github.com/gofiber/fiber/v2"

	"eql/configs"
	"eql/internal/app/repositories"
	"eql/internal/app/usecases/users"
	"eql/internal/handlers/middleware"
	"eql/internal/infrastructures/database"
)

// func SetupRouter(app *fiber.App) {
// 	//สร้างกลุ่มของ Route เพื่อแยกกลุ่มของ API Endpoint ตามแต่ละส่วนในโครงสร้าง
// 	api := app.Group("/api")

// 	UsersRouter(api.Group("/users"))
// 	LogRouter(api.Group("/logs"))

// }

func SetupRouter(app *fiber.App, c *fiber.Ctx) {
	conFigConst := *configs.CF                  //config ต่างๆของ API
	mainDatabase := database.GetMainDatabase(c) //เชื่อมต่อฐานข้อมูลหลัก

	//กำหนดกลุ่ม Repo ที่ต้องใช้
	userRepo := repositories.NewUserRepository(mainDatabase)

	//กำหนด Service

	//กำหนด Endpoint
	userEndPoint := users.NewEndpoint(users.NewService(userRepo, conFigConst)) //Users

	//สร้างกลุ่มของ Route เพื่อแยกกลุ่มของ API Endpoint ตามแต่ละส่วนในโครงสร้าง
	api := app.Group("/api")

	authRoute := api.Group("/auth")
	authRoute.Post("/login", userEndPoint.Login)

	userRoute := api.Group("/users", middleware.JWTMiddleware())
	userRoute.Get("/users2", userEndPoint.GetUserByID)
	userRoute.Get("/getuser", userEndPoint.GetUser)

	LogRouter(api.Group("/logs"))
}
