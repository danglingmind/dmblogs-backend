package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" gorm:"not null;"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	Draft       bool      `json:"draft"`
	UserId      uuid.UUID `json:"user_id"`
}

// Bussiness Services for entity
func (b *Blog) PreSaveValidate() error {
	if b.Title == "" {
		return fmt.Errorf("empty title is not allowed")
	}
	return nil
}
