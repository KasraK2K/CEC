package variant_local

import (
	"time"

	"gorm.io/gorm"
)

type VariantLocal struct {
	ID               uint           `json:"id"                bson:"id"               gorm:"type:uint;primaryKey;<-:false"`
	ModelID          uint           `json:"model_id"          bson:"model_id"         gorm:"type:uint;foreignKey"              validate:"required"`
	Region           string         `json:"region"            bson:"region"           gorm:"type:string;not null;"             validate:"required"`
	Price            int64          `json:"price"             bson:"price"            gorm:"type:bigint"`
	AnnualTax        int64          `json:"annual_tax"        bson:"annual_tax"       gorm:"type:bigint"`
	CongestionCharge int64          `json:"congestion_charge" bson:"congestionCharge" gorm:"type:bigint"`
	InsuranceGroup   int8           `json:"insurance_group"   bson:"insurance_group"  gorm:"type:int8"`
	Availability     int8           `json:"availability"      bson:"availability"     gorm:"type:int8"`
	ShowingRange     int8           `json:"showing_range"     bson:"showing_range"    gorm:"type:int"`
	IsArchive        bool           `json:"is_archive"        bson:"is_archive"       gorm:"default:false"`
	CreatedAt        time.Time      `json:"created_at"        bson:"created_at"       gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt        time.Time      `json:"updated_at"        bson:"updated_at"       gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt        gorm.DeletedAt `json:"archive_at"        bson:"archive_at"       gorm:"type:timestamptz;index"`
}

type VariantLocalFilter struct {
	ID      uint   `json:"id,omitempty"               bson:"id,omitempty"               validate:"omitempty"`
	ModelID uint   `json:"model_id,omitempty"         bson:"model_id,omitempty"         validate:"omitempty"`
	Region  string `json:"region,omitempty"           bson:"region,omitempty"           validate:"omitempty"`
}

type VariantLocalUpdate struct {
	ID               uint   `json:"id"                bson:"id"`
	ModelID          uint   `json:"model_id"          bson:"model_id"`
	Region           string `json:"region"            bson:"region"`
	Price            int64  `json:"price"             bson:"price"`
	AnnualTax        int64  `json:"annual_tax"        bson:"annual_tax"`
	CongestionCharge int64  `json:"congestion_charge" bson:"congestionCharge"`
	InsuranceGroup   int8   `json:"insurance_group"   bson:"insurance_group"`
	Availability     int8   `json:"availability"      bson:"availability"`
	ShowingRange     int8   `json:"showing_range"     bson:"showing_range"`
}
