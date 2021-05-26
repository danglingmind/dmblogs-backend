package repository

import "danglingmind.com/ddd/domain/entity"

type BlogTagRepository interface {
	Save(blogTag entity.BlogTag) error
	Delete(id uint64) error
	DeleteByBlogId(ids []uint64) error
	DeleteByTagId(ids []uint64) error
	GetByBlogId(id uint64) ([]entity.BlogTag, error)
	GetByTagId(id uint64) ([]entity.BlogTag, error)
}
