package repository

import (
	"danglingmind.com/ddd/domain/entity"
	"github.com/google/uuid"
)

type BlogRepository interface {
	Save(blog entity.Blog, userid uuid.UUID) (*entity.Blog, error)
	Delete(id uint64) error
	GetBlogById(id uint64) (*entity.Blog, error)
	GetBlogsByUserId(userid uuid.UUID) ([]entity.Blog, error)
	GetBlogs(limit, offset int) ([]entity.Blog, error)
	GetBlogsByIds(blogIds []uint64, limit, offset int) ([]entity.Blog, error)
}
