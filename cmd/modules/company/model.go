package company

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	CompanyID uint           `json:"company_id" bson:"company_id" gorm:"type:uint;primaryKey;<-:false"`
	Brand     string         `json:"brand"      bson:"brand"      gorm:"type:string;unique;not null;"      validate:"required"`
	IsArchive bool           `json:"is_archive" bson:"is_archive" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at" gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at" gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt gorm.DeletedAt `json:"archive_at" bson:"archive_at" gorm:"type:timestamptz;index"`
}

type CompanyFilter struct {
	CompanyID uint   `json:"company_id,omitempty" bson:"company_id,omitempty" validate:"omitempty"`
	Brand     string `json:"brand,omitempty"      bson:"brand,omitempty"      validate:"omitempty"`
}

type CompanyUpdate struct {
	CompanyID uint   `json:"company_id"     bson:"company_id"`
	Brand     string `json:"brand"          bson:"brand"`
}
