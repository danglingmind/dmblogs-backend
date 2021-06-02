package application

import (
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
)

type BlogTagAppInterface interface {
	Save(blogTag entity.BlogTag) error
	Delete(id uint64) error
	GetByBlogId(id uint64) ([]entity.BlogTag, error)
	GetByTagId(id uint64) ([]entity.BlogTag, error)
	DeleteByBlogId(ids []uint64) error
	DeleteByTagId(ids []uint64) error
}

type BlogTagApp struct {
	bt repository.BlogTagRepository
}

var _ repository.BlogTagRepository = &BlogTagApp{}

func (b *BlogTagApp) Save(blogTag entity.BlogTag) error {
	return b.bt.Save(blogTag)
}

func (b *BlogTagApp) Delete(id uint64) error {
	return b.bt.Delete(id)
}
func (b *BlogTagApp) GetByBlogId(id uint64) ([]entity.BlogTag, error) {
	return b.bt.GetByBlogId(id)
}

func (b *BlogTagApp) GetByTagId(id uint64) ([]entity.BlogTag, error) {
	return b.bt.GetByTagId(id)
}

func (b *BlogTagApp) DeleteByBlogId(ids []uint64) error {
	return b.bt.DeleteByBlogId(ids)
}

func (b *BlogTagApp) DeleteByTagId(ids []uint64) error {
	return b.bt.DeleteByTagId(ids)
}
