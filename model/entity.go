package model

type Bass struct {
	FishCode string `gorm:"column:fish_code"`
	FishName string `gorm:"column:fish_name"`
}
