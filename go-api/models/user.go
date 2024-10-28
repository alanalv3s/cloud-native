package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User model represents a user in the database
type User struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string    `json:"name" binding:"required"`
	Email string    `json:"email" binding:"required,email" gorm:"unique"`
}
