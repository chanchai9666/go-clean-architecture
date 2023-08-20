package router

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"eql/internal/app/entities/models"
)

type User struct {
	ID       int    `json:"id" form:"id" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
}

func UsersRouter(router fiber.Router, c *fiber.Ctx) {

	router.Get("/alluser", func(c *fiber.Ctx) error {
		db1 := c.Locals("main").(*gorm.DB)
		user := []models.User{}
		err := db1.Find(&user).Error
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(user)
	})

	// Define the routes for /users using the provided router
	router.Get("/list", func(c *fiber.Ctx) error {
		fmt.Println("Start")
		time.Sleep(3 * time.Second)
		fmt.Println("End")
		return c.SendString("I'm a GET USER request!")
	})

	// @Summary      Show an account
	// @Description  Get account information by ID
	// @Tags         accounts
	// @Accept       json
	// @Produce      json
	// @Param        id   path      int    true    "Account ID"
	// @Success      200  {object}  AccountResponse
	// @Failure      400  {object}  ErrorResponse
	// @Failure      404  {object}  ErrorResponse
	// @Router       /accounts/{id} [get]
	router.Get("/accounts/{id}", func(c *fiber.Ctx) error {
		// โค้ดที่ทำงานเมื่อมีคำขอเข้ามายัง API endpoint
		return nil
	})

	// router.Get("sss", func(c *fiber.Ctx) error {
	// 	user := []models.User{}
	// 	err := db1.First(&user).Error
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return c.JSON(user)
	// })

	// var user User

	// Use mapToStruct to map and validate input
	// if err := handlers.MapToStruct(c, &user, func(c *fiber.Ctx, input interface{}) error {
	// 	// Call Axc function with the mapped and validated user data
	// 	result, err := Axc(input.(*User))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Return the result as JSON response
	// 	return c.JSON(result)
	// }); err != nil {
	// 	// Handle the error returned from mapToStruct
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	// }

	// 	return nil
	// })
}

func Axc(rq *User) (*User, error) {
	return rq, nil
}
