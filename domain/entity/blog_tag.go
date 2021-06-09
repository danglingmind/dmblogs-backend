package entity

import "time"

type BlogTag struct {
	ID      uint64    `json:"id" gorm:"primary_key;auto_increment"`
	BlogId  uint64    `json:"blog_id"`
	TagId   uint64    `json:"tag_id"`
	Created time.Time `json:"created" gorm:"default:CURRENT_TIMESTAMP"`
	Active  bool      `json:"active" gorm:"default:1;"`
}

func (b *BlogTag) PreSave() error {
	return nil
}
