package models

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     string
	Password  string
	Status    int
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
