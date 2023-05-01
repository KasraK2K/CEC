package models

import (
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID                            uint           `json:"id"                               bson:"id"                               gorm:"type:uint;primaryKey;<-:false"`
	CompanyID                     uint           `json:"company_id"                       bson:"company_id"                       gorm:"type:uint;"                        validate:"required"`
	ModelID                       uint           `json:"model_id"                         bson:"model_id"                         gorm:"type:uint;unique"                  validate:"required"`
	VariantName                   string         `json:"variant_name"                     bson:"variant_name"                     gorm:"type:string;not null;"`
	Seat                          int            `json:"seat"                             bson:"seat"                             gorm:"type:int"`
	CarBody                       string         `json:"car_body"                         bson:"car_body"                         gorm:"type:string;not null;"`
	Type                          string         `json:"type"                             bson:"type"                             gorm:"type:string;not null;"`
	RealRange                     int            `json:"real_range"                       bson:"real_range"                       gorm:"type:int"`
	ReleaseYear                   int            `json:"release_year"                     bson:"release_year"                     gorm:"type:int"`
	RangeCityCold                 int            `json:"range_city_cold"                  bson:"range_city_cold"                  gorm:"type:int"`
	RangeCityMild                 int            `json:"range_city_mild"                  bson:"range_city_mild"                  gorm:"type:int"`
	RangeCombinedCold             int            `json:"range_combined_cold"              bson:"range_combined_cold"              gorm:"type:int"`
	RangeCombinedMild             int            `json:"range_combined_mild"              bson:"range_combined_mild"              gorm:"type:int"`
	RangeHighwayCold              int            `json:"range_highway_cold"               bson:"range_highway_cold"               gorm:"type:int"`
	RangeHighwayMild              int            `json:"range_highway_mild"               bson:"range_highway_mild"               gorm:"type:int"`
	BatteryCapacity               int            `json:"battery_capacity"                 bson:"battery_capacity"                 gorm:"type:int"`
	BatteryUseable                int            `json:"battery_useable"                  bson:"battery_useable"                  gorm:"type:int"`
	ChargeTimeMinuteRegular       int            `json:"charge_time_minute_regular"       bson:"charge_time_minute_regular"       gorm:"type:int"`
	ChargeSpeedRegular            int            `json:"charge_speed_regular"             bson:"charge_speed_regular"             gorm:"type:int"`
	ChargePowerMaxRegular         int            `json:"charge_power_max_regular"         bson:"charge_power_max_regular"         gorm:"type:int"`
	ChargePortCharacterRegular    string         `json:"charge_port_character_regular"    bson:"charge_port_character_regular"    gorm:"type:string;not null;"`
	ChargePortLocationRegular     string         `json:"charge_port_location_regular"     bson:"charge_port_location_regular"     gorm:"type:string;not null;"`
	ChargeTimeMinuteFastcharge    int            `json:"charge_time_minute_fastcharge"    bson:"charge_time_minute_fastcharge"    gorm:"type:int"`
	ChargeSpeedFastcharge         int            `json:"charge_speed_fastcharge"          bson:"charge_speed_fastcharge"          gorm:"type:int"`
	ChargePowerMaxFastcharge      int            `json:"charge_power_max_fastcharge"      bson:"charge_power_max_fastcharge"      gorm:"type:int"`
	ChargePortCharacterFastcharge string         `json:"charge_port_character_fastcharge" bson:"charge_port_character_fastcharge" gorm:"type:string;not null;"`
	ChargePortLocationFastcharge  string         `json:"charge_port_location_fastcharge"  bson:"charge_port_location_fastcharge"  gorm:"type:string;not null;"`
	ConsumptionCityCold           int            `json:"consumption_city_cold"            bson:"consumption_city_cold"            gorm:"type:int"`
	ConsumptionHighwayCold        int            `json:"consumption_highway_cold"         bson:"consumption_highway_cold"         gorm:"type:int"`
	ConsumptionCombinedCold       int            `json:"consumption_combined_cold"        bson:"consumption_combined_cold"        gorm:"type:int"`
	ConsumptionCityMild           int            `json:"consumption_city_mild"            bson:"consumption_city_mild"            gorm:"type:int"`
	ConsumptionHighwayMild        int            `json:"consumption_highway_mild"         bson:"consumption_highway_mild"         gorm:"type:int"`
	ConsumptionCombinedMild       int            `json:"consumption_combined_mild"        bson:"consumption_combined_mild"        gorm:"type:int"`
	AccelerationReal              int            `json:"acceleration"                     bson:"acceleration"                     gorm:"type:int"`
	TopSpeed                      int            `json:"top_speed"                        bson:"top_speed"                        gorm:"type:int"`
	TotalPowerKw                  int            `json:"total_power_kw"                   bson:"total_power_kw"                   gorm:"type:int"`
	Co2Emissions                  int            `json:"co2_emissions"                    bson:"co2_emissions"                    gorm:"type:int"`
	Drive                         string         `json:"drive"                            bson:"drive"                            gorm:"type:string;not null;"`
	Length                        int            `json:"length"                           bson:"length"                           gorm:"type:int"`
	Width                         int            `json:"width"                            bson:"width"                            gorm:"type:int"`
	WidthWithMirror               int            `json:"width_with_mirror"                bson:"width_with_mirror"                gorm:"type:int"`
	Height                        int            `json:"height"                           bson:"height"                           gorm:"type:int"`
	Wheelbase                     int            `json:"wheelbase"                        bson:"wheelbase"                        gorm:"type:int"`
	WeightUnladen                 int            `json:"weight_unladen"                   bson:"weight_unladen"                   gorm:"type:int"`
	WeightGrossVehicle            int            `json:"weight_gross_vehicle"             bson:"weight_gross_vehicle"             gorm:"type:int"`
	MaxPayload                    int            `json:"max_payload"                      bson:"max_payload"                      gorm:"type:int"`
	TowingWeightUnBraked          int            `json:"towing_weight_un_braked"          bson:"towing_weight_un_braked"          gorm:"type:int"`
	TowingWeightBraked            int            `json:"towing_weight_braked"             bson:"towing_weight_braked"             gorm:"type:int"`
	VerticalLoadMax               int            `json:"vertical_load_max"                bson:"vertical_load_max"                gorm:"type:int"`
	CargoVolume                   int            `json:"cargo_volume"                     bson:"cargo_volume"                     gorm:"type:int"`
	CargoVolumeMax                int            `json:"cargo_volume_max"                 bson:"cargo_volume_max"                 gorm:"type:int"`
	CargoVolumeFrunk              int            `json:"cargo_volume_frunk"               bson:"cargo_volume_frunk"               gorm:"type:int"`
	RoofLoad                      int            `json:"roof_load"                        bson:"roof_load"                        gorm:"type:int"`
	RoofRails                     int            `json:"roof_rails"                       bson:"roof_rails"                       gorm:"type:int"`
	Isofix                        int            `json:"isofix"                           bson:"isofix"                           gorm:"type:int"`
	TowHitchPossible              int            `json:"tow_hitch_possible"               bson:"tow_hitch_possible"               gorm:"type:int"`
	IsArchive                     bool           `json:"is_archive"                       bson:"is_archive"                       gorm:"default:false"`
	CreatedAt                     time.Time      `json:"created_at"                       bson:"created_at"                       gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt                     time.Time      `json:"updated_at"                       bson:"updated_at"                       gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt                     gorm.DeletedAt `json:"archive_at"                       bson:"archive_at"                       gorm:"type:timestamptz;index"`
}

type FindVariant struct {
	Variant
	Model   *Model   `json:"model"   bson:"model"   gorm:"foreignKey:ModelID"`   // BelongsTo Model
	Company *Company `json:"company" bson:"company" gorm:"foreignKey:CompanyID"` // BelongsTo Company
}

type VariantFilter struct {
	ID        uint `json:"id,omitempty"         bson:"id,omitempty"         validate:"omitempty"`
	ModelID   uint `json:"model_id,omitempty"   bson:"model_id,omitempty"   validate:"omitempty"`
	CompanyID uint `json:"company_id,omitempty" bson:"company_id,omitempty" validate:"omitempty"`
}

type VariantUpdate struct {
	ID                            uint   `json:"id"                               bson:"id"`
	CompanyID                     uint   `json:"company_id"                       bson:"company_id"`
	ModelID                       uint   `json:"model_id"                         bson:"model_id"`
	VariantName                   string `json:"variant_name"                     bson:"variant_name"`
	Seat                          int    `json:"seat"                             bson:"seat"`
	CarBody                       string `json:"car_body"                         bson:"car_body"`
	Type                          string `json:"type"                             bson:"type"`
	RealRange                     int    `json:"real_range"                       bson:"real_range"`
	ReleaseYear                   int    `json:"release_year"                     bson:"release_year"`
	RangeCityCold                 int    `json:"range_city_cold"                  bson:"range_city_cold"`
	RangeCityMild                 int    `json:"range_city_mild"                  bson:"range_city_mild"`
	RangeCombinedCold             int    `json:"range_combined_cold"              bson:"range_combined_cold"`
	RangeCombinedMild             int    `json:"range_combined_mild"              bson:"range_combined_mild"`
	RangeHighwayCold              int    `json:"range_highway_cold"               bson:"range_highway_cold"`
	RangeHighwayMild              int    `json:"range_highway_mild"               bson:"range_highway_mild"`
	BatteryCapacity               int    `json:"battery_capacity"                 bson:"battery_capacity"`
	BatteryUseable                int    `json:"battery_useable"                  bson:"battery_useable"`
	ChargeTimeMinuteRegular       int    `json:"charge_time_minute_regular"       bson:"charge_time_minute_regular"`
	ChargeSpeedRegular            int    `json:"charge_speed_regular"             bson:"charge_speed_regular"`
	ChargePowerMaxRegular         int    `json:"charge_power_max_regular"         bson:"charge_power_max_regular"`
	ChargePortCharacterRegular    string `json:"charge_port_character_regular"    bson:"charge_port_character_regular"`
	ChargePortLocationRegular     string `json:"charge_port_location_regular"     bson:"charge_port_location_regular"`
	ChargeTimeMinuteFastcharge    int    `json:"charge_time_minute_fastcharge"    bson:"charge_time_minute_fastcharge"`
	ChargeSpeedFastcharge         int    `json:"charge_speed_fastcharge"          bson:"charge_speed_fastcharge"`
	ChargePowerMaxFastcharge      int    `json:"charge_power_max_fastcharge"      bson:"charge_power_max_fastcharge"`
	ChargePortCharacterFastcharge string `json:"charge_port_character_fastcharge" bson:"charge_port_character_fastcharge"`
	ChargePortLocationFastcharge  string `json:"charge_port_location_fastcharge"  bson:"charge_port_location_fastcharge"`
	ConsumptionCityCold           int    `json:"consumption_city_cold"            bson:"consumption_city_cold"`
	ConsumptionHighwayCold        int    `json:"consumption_highway_cold"         bson:"consumption_highway_cold"`
	ConsumptionCombinedCold       int    `json:"consumption_combined_cold"        bson:"consumption_combined_cold"`
	ConsumptionCityMild           int    `json:"consumption_city_mild"            bson:"consumption_city_mild"`
	ConsumptionHighwayMild        int    `json:"consumption_highway_mild"         bson:"consumption_highway_mild"`
	ConsumptionCombinedMild       int    `json:"consumption_combined_mild"        bson:"consumption_combined_mild"`
	AccelerationReal              int    `json:"acceleration"                bson:"acceleration"`
	TopSpeed                      int    `json:"top_speed"                        bson:"top_speed"`
	TotalPowerKw                  int    `json:"total_power_kw"                   bson:"total_power_kw"`
	Co2Emissions                  int    `json:"co2_emissions"                    bson:"co2_emissions"`
	Drive                         string `json:"drive"                            bson:"drive"`
	Length                        int    `json:"length"                           bson:"length"`
	Width                         int    `json:"width"                            bson:"width"`
	WidthWithMirror               int    `json:"width_with_mirror"                bson:"width_with_mirror"`
	Height                        int    `json:"height"                           bson:"height"`
	Wheelbase                     int    `json:"wheelbase"                        bson:"wheelbase"`
	WeightUnladen                 int    `json:"weight_unladen"                   bson:"weight_unladen"`
	WeightGrossVehicle            int    `json:"weight_gross_vehicle"             bson:"weight_gross_vehicle"`
	MaxPayload                    int    `json:"max_payload"                      bson:"max_payload"`
	TowingWeightUnBraked          int    `json:"towing_weight_un_braked"          bson:"towing_weight_un_braked"`
	TowingWeightBraked            int    `json:"towing_weight_braked"             bson:"towing_weight_braked"`
	VerticalLoadMax               int    `json:"vertical_load_max"                bson:"vertical_load_max"`
	CargoVolume                   int    `json:"cargo_volume"                     bson:"cargo_volume"`
	CargoVolumeMax                int    `json:"cargo_volume_max"                 bson:"cargo_volume_max"`
	CargoVolumeFrunk              int    `json:"cargo_volume_frunk"               bson:"cargo_volume_frunk"`
	RoofLoad                      int    `json:"roof_load"                        bson:"roof_load"`
	RoofRails                     int    `json:"roof_rails"                       bson:"roof_rails"`
	Isofix                        int    `json:"isofix"                           bson:"isofix"`
	TowHitchPossible              int    `json:"tow_hitch_possible"               bson:"tow_hitch_possible"`
}
