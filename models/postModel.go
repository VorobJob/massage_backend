package models

import (
	"time"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model

	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"size:40; not null"`
	PhoneNum       string    `gorm:"size:20; not null"`
	StartTime      time.Time `gorm:"default:17:00:00"`
	EndTime        time.Time `gorm:"default:05:00:00"`
	DaysOfWeek     [7]int    `gorm:"type:integer[]"`
	InhousePrices  [3]int    `gorm:"type:integer[]; not null"`
	OutPrices      [2]int    `gorm:"type:integer[]; not null"`
	ClientsSex     int       `gorm:"type:integer; not null"`
	IsMinAge       bool      `gorm:"default:true"`
	City           int       `gorm:"type:integer; not null"`
	MetroStations  []int     `gorm:"type:integer[]; not null"`
	Address        string    `gorm:"size:255; not null"`
	AddressComment string    `gorm:"size:300"`
	Age            int       `gorm:"not null"`
	Education      string    `gorm:"size:150"`
	Experience     int
	Bio            string `gorm:"size:1000"`
	Created        time.Time
	LastUpdate     time.Time
	ServiceTypes   []int    `gorm:"type:integer[]; not null"`
	Photos         []string `gorm:"type:varchar[]; not null"`
	IsSalon        bool     `gorm:"default:false; not null"`
}
