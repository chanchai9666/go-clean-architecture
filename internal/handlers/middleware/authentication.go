package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"eql/internal/app/entities/models"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ดึง Token จาก Header "Authorization"
		tokenString := c.Get("Authorization")

		// ตรวจสอบว่า Token ถูกส่งมาหรือไม่
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// ตรวจสอบว่า Token ถูกตรวจสอบและถูกเซ็นด้วยคีย์ลับที่ถูกต้องหรือไม่
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1) // ตัดคำว่า "Bearer " ออก
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // ใส่คีย์ลับของคุณที่นี่
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// ตรวจสอบว่า Token ถูกเซ็นด้วยวิธีการที่ถูกต้อง
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// ตรวจสอบข้อมูลที่ถูกเก็บใน Token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// ตรวจสอบวันหมดอายุของ Token
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token has expired",
			})
		}

		// ตรวจสอบข้อมูลผู้ใช้จาก Token
		userID := int(claims["user_id"].(float64))
		userName := claims["user_name"].(string)

		// ทำสิ่งที่คุณต้องการกับข้อมูลผู้ใช้ (user)
		// เช่น การใช้ข้อมูลผู้ใช้ในระบบหรือการตรวจสอบสิทธิ์
		user := models.User{
			ID:       userID,
			Username: userName,
		}

		// ส่งข้อมูลผู้ใช้ไปยัง Context ของ Fiber
		c.Locals("user", user)

		// ดำเนินการถัดไปในการร้องขอ
		return c.Next()
	}
}
