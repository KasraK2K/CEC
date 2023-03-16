package user

import (
	"CEC/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"type:uint;primaryKey;<-:false"`
	FirstName string         `gorm:"type:string; not null;" json:"first_name"`
	LastName  string         `gorm:"type:string;" json:"last_name"`
	UserName  string         `gorm:"type:string;unique;<-:create" json:"user_name" validate:"required,min=3,max=15"`
	Password  string         `gorm:"type:string;check:length(password) >= 8" json:"password" validate:"required,min=8,max=32"`
	Email     string         `gorm:"type:string;unique;not null;" json:"email" validate:"required,email,min=6,max=32"`
	Phone     string         `gorm:"type:string;unique;" json:"phone"`
	CreatedAt time.Time      `gorm:"type:datetime;autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:datetime;autoUpdateTime;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime;index" json:"deleted_at"`
}

func (user *User) Validate() helper.ErrorResponse {
	return helper.Validator(&user)
}
