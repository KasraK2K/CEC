package company

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID        uint           `json:"id"         bson:"id"         gorm:"type:uint;primaryKey;<-:false"`
	Brand     string         `json:"brand"      bson:"brand"      gorm:"type:string;unique;not null;"      validate:"required"`
	IsArchive bool           `json:"is_archive" bson:"is_archive" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at" gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at" gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt gorm.DeletedAt `json:"archive_at" bson:"archive_at" gorm:"type:timestamptz;index"`
}

type CompanyFilter struct {
	ID    uint   `json:"id,omitempty"    bson:"id,omitempty"    validate:"omitempty"`
	Brand string `json:"brand,omitempty" bson:"brand,omitempty" validate:"omitempty"`
}

type CompanyUpdate struct {
	ID    uint   `json:"id"    bson:"id"`
	Brand string `json:"brand" bson:"brand"`
}
