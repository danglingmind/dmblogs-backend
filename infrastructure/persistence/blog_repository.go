package persistence

import (
	"fmt"
	"strings"

	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BlogRepo struct {
	db *gorm.DB
}

var _ repository.BlogRepository = &BlogRepo{}

func NewBlogRepository(db *gorm.DB) *BlogRepo {
	return &BlogRepo{
		db: db,
	}
}

func (b *BlogRepo) Save(blog entity.Blog, userid uuid.UUID) (*entity.Blog, error) {
	err := blog.PreSaveValidate()
	if err != nil {
		return nil, err
	}

	err = b.db.Debug().
		Table("blogs").
		Create(&blog).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, fmt.Errorf("please enter a unique title")
		}
		return nil, err
	}

	return &blog, err
}

func (b *BlogRepo) Delete(id uint64) error {
	var blog entity.Blog
	blog.ID = id
	return b.db.Debug().Delete(&blog).Error

}

func (b *BlogRepo) GetBlogById(id uint64) (*entity.Blog, error) {
	var blog entity.Blog
	err := b.db.Debug().
		Table("blogs").
		Where("id = ?", id).
		Find(&blog).
		Error
	if err != nil {
		return nil, err
	}
	return &blog, err
}

func (b *BlogRepo) GetBlogsByUserId(userid uuid.UUID) ([]entity.Blog, error) {
	var blogs []entity.Blog
	err := b.db.Debug().
		Table("blogs").
		Where("user_id = ?", userid).
		Find(&blogs).
		Error
	return blogs, err
}

func (b *BlogRepo) GetBlogs(limit, offset int) ([]entity.Blog, error) {
	var blogs []entity.Blog
	err := b.db.Debug().
		Table("blogs").
		Limit(limit).
		Offset(offset).
		Find(&blogs).
		Error
	return blogs, err
}

func (b *BlogRepo) GetBlogsByIds(blogIds []uint64, limit, offset int) ([]entity.Blog, error) {
	var blogs []entity.Blog
	err := b.db.Debug().
		Table("blogs").
		Where("id IN ?", blogIds).
		Limit(limit).
		Offset(offset).
		Find(&blogs).
		Error
	return blogs, err
}
