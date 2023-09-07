package users

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"eql/configs"
	"eql/internal/app/entities/models"
	"eql/internal/app/entities/schema"
	"eql/internal/app/repositories"
	"eql/internal/app/service"
)

type Service interface {
	GetUserAll(c *fiber.Ctx) ([]models.User, error)
	GetUser(c *fiber.Ctx, req *schema.UserRequest) ([]models.User, error)
	GetUserByUserName(c *fiber.Ctx, req *schema.LoginReq) (*schema.LoginResp, error)
}

type userService struct {
	repo      repositories.UserRepository
	appConfig configs.Config
}

func NewService(repo repositories.UserRepository, appConfig configs.Config) Service {
	return &userService{
		repo:      repo,
		appConfig: appConfig,
	}
}

func (s *userService) GetUserAll(c *fiber.Ctx) ([]models.User, error) {

	user, err := s.repo.GetUserData()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUser(c *fiber.Ctx, req *schema.UserRequest) ([]models.User, error) {
	user, err := s.repo.GetUser(&models.User{
		Username: req.Username,
		ID:       req.ID,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUserName(c *fiber.Ctx, req *schema.LoginReq) (*schema.LoginResp, error) {
	data, err := s.repo.GetUserByUserName(req.Username)
	if err != nil {
		return nil, err
	}
	// ตรวจสอบรหัสผ่าน
	if !service.CheckPassword(req.Password, data.Password) {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "username or password is incorrect.")
	}

	// สร้าง JWT Token
	token, err := createJWTToken(data)
	if err != nil {
		return nil, err
	}

	rReturn := schema.LoginResp{
		ID:          data.ID,
		Username:    data.Username,
		Email:       data.Email,
		AccessToken: token,
	}

	return &rReturn, nil
}

func createJWTToken(user *models.User) (string, error) {
	// สร้างข้อมูลที่จะใช้ในการสร้าง Token
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // กำหนดวันหมดอายุให้ Token
	}

	// สร้าง Token ด้วยคีย์ลับ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-secret-key")) // ใส่คีย์ลับของคุณที่นี่

	if err != nil {
		return "", err
	}

	// ส่ง Token ในรูปแบบ Bearer Token
	return "Bearer " + tokenString, nil
}
