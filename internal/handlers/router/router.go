// handlers/router/router.go
package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	//สร้างกลุ่มของ Route เพื่อแยกกลุ่มของ API Endpoint ตามแต่ละส่วนในโครงสร้าง
	api := app.Group("/api")

	UsersRouter(api.Group("/users"))
	LogRouter(api.Group("/logs"))

}
