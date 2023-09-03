// handlers/router/router.go
package router

import (
	"github.com/gofiber/fiber/v2"

	"eql/internal/app/repositories"
	"eql/internal/app/usecases/users"
	"eql/internal/infrastructures/database"
)

// func SetupRouter(app *fiber.App) {
// 	//สร้างกลุ่มของ Route เพื่อแยกกลุ่มของ API Endpoint ตามแต่ละส่วนในโครงสร้าง
// 	api := app.Group("/api")

// 	UsersRouter(api.Group("/users"))
// 	LogRouter(api.Group("/logs"))

// }

func SetupRouter(app *fiber.App, c *fiber.Ctx) {
	//กำหนดกลุ่ม Repo ที่ต้องใช้
	userEndPoint := users.NewEndpoint(users.NewService(repositories.NewUserRepository(database.GetMainDatabase(c))))

	//สร้างกลุ่มของ Route เพื่อแยกกลุ่มของ API Endpoint ตามแต่ละส่วนในโครงสร้าง
	api := app.Group("/api")

	userRouter := api.Group("/users")
	userRouter.Get("/users2", userEndPoint.GetUserByID)
	userRouter.Get("/getuser", userEndPoint.GetUser)

	LogRouter(api.Group("/logs"))
}
