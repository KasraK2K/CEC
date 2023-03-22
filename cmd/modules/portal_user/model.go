package portal_user

import (
	"time"

	"gorm.io/gorm"
)

type PortalUser struct {
	gorm.Model
	ID            uint           `json:"id"                 bson:"id"                 gorm:"type:uint;primaryKey;<-:false"`
	Email         string         `json:"email"              bson:"email"              gorm:"type:string;unique;not null;"                        validate:"required,email,min=6,max=32"`
	Password      string         `json:"password,omitempty" bson:"password,omitempty" gorm:"type:string;check:length(password) >= 8"             validate:"required,min=8,max=32"`
	ContactNumber string         `json:"contact_number"     bson:"contact_number"     gorm:"type:string;"`
	FirstName     string         `json:"first_name"         bson:"first_name"         gorm:"type:string;"`
	Surname       string         `json:"surname"            bson:"surname"            gorm:"type:string;"`
	Gender        uint8          `json:"gender"             bson:"gender"             gorm:"type:int8;default:2"`
	IsActive      bool           `json:"is_active"          bson:"is_active"          gorm:"default:true"`
	IsAdmin       bool           `json:"is_admin"           bson:"is_admin"           gorm:"default:false"`
	IsArchive     bool           `json:"is_archive"         bson:"is_archive"         gorm:"default:false"`
	CreatedAt     time.Time      `json:"created_at"         bson:"created_at"         gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt     time.Time      `json:"updated_at"         bson:"updated_at"         gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt     gorm.DeletedAt `json:"archive_at"         bson:"archive_at"         gorm:"type:timestamptz;index"`
}

type PortalUserFilter struct {
	ID        uint   `json:"id,omitempty"         bson:"id,omitempty"`
	Email     string `json:"email,omitempty"      bson:"email,omitempty"`
	Gender    uint8  `json:"gender,omitempty"     bson:"gender,omitempty"`
	IsActive  bool   `json:"is_active,omitempty"  bson:"is_active,omitempty"`
	IsAdmin   bool   `json:"is_admin,omitempty"   bson:"is_admin,omitempty"`
	IsArchive bool   `json:"is_archive,omitempty" bson:"is_archive,omitempty"`
}

type PortalUserUpdate struct {
	gorm.Model
	ID            uint   `json:"id"             bson:"id"`
	Email         string `json:"email"          bson:"email"`
	Password      string `json:"password"       bson:"password"`
	ContactNumber string `json:"contact_number" bson:"contact_number"`
	FirstName     string `json:"first_name"     bson:"first_name"`
	Surname       string `json:"surname"        bson:"surname"`
	Gender        uint8  `json:"gender"         bson:"gender"`
	IsActive      bool   `json:"is_active"      bson:"is_active"`
	IsAdmin       bool   `json:"is_admin"       bson:"is_admin"`
	IsArchive     bool   `json:"is_archive"     bson:"is_archive"`
}

type PortalUserLoginPayload struct {
	Email    string `json:"email"    bson:"email"    validate:"required,email,min=6,max=32"`
	Password string `json:"password" bson:"password" validate:"required,min=8,max=32"`
	Platform uint8  `json:"platform" bson:"platform" validate:"required,oneof=1 2"`
}
