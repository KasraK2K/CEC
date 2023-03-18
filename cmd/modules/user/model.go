package user

import (
	"time"

	"gorm.io/gorm"

	"app/pkg/helper"
)

type User struct {
	ID        uint           `json:"id"         bson:"id"         gorm:"type:uint;primaryKey;<-:false"`
	FirstName string         `json:"first_name" bson:"first_name" gorm:"type:string;not null;"`
	Surname   string         `json:"surname"    bson:"surname"    gorm:"type:string;"`
	UserName  string         `json:"user_name"  bson:"user_name"  gorm:"type:string;unique;<-:create"                        validate:"required,min=3,max=15"`
	Password  string         `json:"password"   bson:"password"   gorm:"type:string;check:length(password) >= 8"             validate:"required,min=8,max=32"`
	Email     string         `json:"email"      bson:"email"      gorm:"type:string;unique;not null;"                        validate:"required,email,min=6,max=32"`
	Phone     string         `json:"phone"      bson:"phone"      gorm:"type:string;unique;"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at" gorm:"type:timestamptz;column:created_at;autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at" gorm:"type:timestamptz;column:updated_at;autoUpdateTime;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" bson:"deleted_at" gorm:"type:timestamptz;column:deleted_at;index"`
}

func (user *User) Validate() helper.ErrorResponse {
	return helper.Validator(user)
}
