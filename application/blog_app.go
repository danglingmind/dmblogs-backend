package application

import (
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
)

type BlogAppInterface interface {
	Save(blog entity.Blog, userid uint64) (*entity.Blog, error)
	Delete(id uint64) error
	GetBlogById(id uint64) (*entity.Blog, error)
	GetBlogs(limit, offset int) ([]entity.Blog, error)
	GetBlogsByIds(blogIds []uint64, limit, offset int) ([]entity.Blog, error)
}

type BlogApp struct {
	bg repository.BlogRepository
}

var _ repository.BlogRepository = &BlogApp{}

func (b *BlogApp) Save(blog entity.Blog, userid uint64) (*entity.Blog, error) {
	return b.bg.Save(blog, userid)
}

func (b *BlogApp) Delete(id uint64) error {
	return b.bg.Delete(id)
}

func (b *BlogApp) GetBlogById(id uint64) (*entity.Blog, error) {
	return b.bg.GetBlogById(id)
}

func (b *BlogApp) GetBlogs(limit, offset int) ([]entity.Blog, error) {
	return b.bg.GetBlogs(limit, offset)
}

func (b *BlogApp) GetBlogsByIds(blogIds []uint64, limit, offset int) ([]entity.Blog, error) {
	return b.bg.GetBlogsByIds(blogIds, limit, offset)
}
