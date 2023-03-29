package model

import (
	"time"

	"gorm.io/gorm"

	"app/cmd/modules/company"
)

type Model struct {
	ModelID   uint           `json:"model_id"   bson:"model_id"   gorm:"type:uint;primaryKey;<-:false"`
	CompanyID uint           `json:"company_id" bson:"company_id" gorm:"type:uint;"                        validate:"required"`
	Model     string         `json:"model"      bson:"model"      gorm:"type:string;unique;not null;"      validate:"required"`
	IsArchive bool           `json:"is_archive" bson:"is_archive" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at" bson:"created_at" gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt time.Time      `json:"updated_at" bson:"updated_at" gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt gorm.DeletedAt `json:"archive_at" bson:"archive_at" gorm:"type:timestamptz;index"`
}

type FindModel struct {
	Model
	Company company.Company `json:"company" bson:"company"`
}

type ModelFilter struct {
	ModelID   uint   `json:"model_id,omitempty"   bson:"model_id,omitempty"   validate:"omitempty"`
	CompanyID uint   `json:"company_id,omitempty" bson:"company_id,omitempty" validate:"omitempty"`
	Model     string `json:"model,omitempty"      bson:"model,omitempty"      validate:"omitempty"`
}

type ModelUpdate struct {
	ModelID   uint   `json:"model_id"       bson:"model_id"`
	CompanyID uint   `json:"company_id"     bson:"company_id"`
	Model     string `json:"model"          bson:"model"`
}
