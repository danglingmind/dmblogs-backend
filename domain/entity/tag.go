package entity

import "time"

type Tag struct {
	ID      uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name    string    `json:"name" gorm:"size:100;not null;"`
	Created time.Time `json:"created" gorm:"default:CURRENT_TIMESTAMP"`
	Active  bool      `json:"active" gorm:"default:1;"`
}

func (t *Tag) PreSave() error {
	return nil
}
