package model

import "time"

type Abstract struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     uint      `json:"created_by"`
	CreatedByName string    `json:"created_by_name"`
	UpdatedAt     time.Time `json:"updated_at"`
	UpdatedBy     uint      `json:"updated_by"`
	UpdatedByName string    `json:"updated_by_name"`
	Active        bool      `json:"active" gorm:"default:true"`
}
