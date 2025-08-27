package models

import "time"

type Users struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Username  string    `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null;comment:密码" json:"password"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}

func (Users) TableName() string {
	return "users"
}
