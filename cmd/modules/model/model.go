package model

import (
	"time"

	"gorm.io/gorm"

	"app/cmd/modules/company"
)

type Model struct {
	ID        uint           `json:"id"         bson:"id"         gorm:"type:uint;primaryKey;<-:false"`
	CompanyID uint           `json:"company_id" bson:"company_id" gorm:"type:uint;foreignKey"              validate:"required"`
	Model     string         `json:"model"      bson:"model"      gorm:"type:string;unique;not null;"      validate:"required"`
	IsArchive bool           `json:"is_archive" bson:"is_archive" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at" gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at" gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt gorm.DeletedAt `json:"archive_at" bson:"archive_at" gorm:"type:timestamptz;index"`
}

type FindModel struct {
	Model
	Company *company.Company `json:"company,omitempty"       bson:"company,omitempty"`
}

type ModelFilter struct {
	ID        uint   `json:"id,omitempty"   bson:"id,omitempty"               validate:"omitempty"`
	CompanyID uint   `json:"company_id,omitempty" bson:"company_id,omitempty" validate:"omitempty"`
	Model     string `json:"model,omitempty"      bson:"model,omitempty"      validate:"omitempty"`
}

type ModelUpdate struct {
	ID        uint   `json:"id"             bson:"id"`
	CompanyID uint   `json:"company_id"     bson:"company_id"`
	Model     string `json:"model"          bson:"model"`
}
