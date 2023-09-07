package repositories

import (
	"gorm.io/gorm"

	"eql/internal/app/entities/models"
)

type UserRepository interface {
	GetUserData() ([]models.User, error)
	GetUser(req *models.User) ([]models.User, error)
	GetUserByUserName(UserName string) (*models.User, error)
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
	// data := []models.User{}
	// for i := 1; i <= 5000; i++ {
	// 	ee := "admin" + fmt.Sprintf("%v", i) + "@admin.com"
	// 	data = append(data, models.User{
	// 		Username: "admin" + fmt.Sprintf("%v", i),
	// 		Password: "1234",
	// 		Email:    ee,
	// 	})
	// }
	// err := r.db.CreateInBatches(data, 100).Error
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(len(data))

	// return nil, nil
}

func (r *userRepository) GetUser(req *models.User) ([]models.User, error) {
	var entities []models.User
	tx := r.db
	if req.Username != "" {
		tx = tx.Where("username=?", req.Username)
	}
	if req.ID > 0 {
		tx = tx.Where("id=?", req.ID)
	}
	if req.Password != "" {
		tx = tx.Where("password=?", req.Password)
	}
	err := tx.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *userRepository) GetUserByUserName(UserName string) (*models.User, error) {
	var entities models.User
	err := r.db.Debug().Where("username=?", UserName).First(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}
