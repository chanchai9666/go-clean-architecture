package repositories

import (
	"gorm.io/gorm"

	"eql/internal/app/entities/models"
	"eql/internal/app/entities/schema"
)

type UserRepository interface {
	GetUserData() ([]models.User, error)
	GetUser(req *schema.UserRequest) ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserData() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUser(req *schema.UserRequest) ([]models.User, error) {

	var user []models.User
	tx := r.db.Debug()
	if req.Username != "" {
		tx = tx.Where("username=?", req.Username)
	}
	err := tx.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
