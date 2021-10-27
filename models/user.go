package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string `gorm:"type:uuid;primaryKey;"`
	FirstName string `gorm:"size:255;unique"`
	LastName  string
	Author    Author    `gorm:"embedded"`
	CreatedAt time.Time `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
	updatedAt time.Time `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL"`
	Books		[]Books `gorm:"many2many:user_books;"`
}

type Author struct {
	Email    string `gorm:"unique"`
	Password string
}

type Books struct {
	ID     string `gorm:"primaryKey"`
	UserId string
	User   User `gorm:"foreignKey:UserId"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.NewString()
	return nil

	// you can also set ID before any User create to table in database by doing this;

	//scope.Statement.SetColumn("ID", uuid.NewString())
	//return nil
}
