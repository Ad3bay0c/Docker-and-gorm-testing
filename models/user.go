package models

import "time"

type User struct {
	ID        string	`gorm:"primaryKey"`
	FirstName string	`gorm:"size:255;unique"`
	LastName  string
	Author    Author		`gorm:"embedded"`
	CreatedAt time.Time	`gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
	updatedAt time.Time	`gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
}

type Author struct {
	Email	string	`gorm:"unique"`
	Password	string
}

type Books struct {
	ID	string	`gorm:"primaryKey"`
	UserId	string
	User	User	`gorm:"foreignKey:UserId"`
}