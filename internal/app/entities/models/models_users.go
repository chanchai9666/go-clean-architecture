package models

import "time"

type User struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;uniqueIndex" json:"username"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Password  string    `gorm:"column:password" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName คืนชื่อของตารางในฐานข้อมูลที่เกี่ยวข้องกับโครงสร้างข้อมูล User
func (User) TableName() string {
	return "user"
}
