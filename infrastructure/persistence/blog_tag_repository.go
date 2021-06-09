package persistence

import (
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
	"github.com/jinzhu/gorm"
)

type BlogTagRepo struct {
	db *gorm.DB
}

func NewBlogTagRepo(db *gorm.DB) *BlogTagRepo {
	return &BlogTagRepo{db: db}
}

var _ repository.BlogTagRepository = &BlogTagRepo{}

func (bt *BlogTagRepo) Save(blogTag entity.BlogTag) error {
	err := blogTag.PreSave()

	if err != nil {
		return err
	}
	return bt.db.Debug().
		Table("blogtags").
		Create(&blogTag).
		Error
}

func (bt *BlogTagRepo) Delete(id uint64) error {
	var blogtag entity.BlogTag
	blogtag.ID = id
	return bt.db.Debug().Table("blogtag").Delete(&blogtag).Error
}

func (bt *BlogTagRepo) DeleteByBlogId(ids []uint64) error {
	return bt.db.
		Debug().
		Table("blogtags").
		Where("blog_id IN ?", ids).
		Delete(entity.BlogTag{}).
		Error
}

func (bt *BlogTagRepo) DeleteByTagId(ids []uint64) error {
	return bt.db.
		Debug().
		Table("blogtags").
		Where("tag_id IN ?", ids).
		Delete(entity.BlogTag{}).
		Error
}

func (bt *BlogTagRepo) GetByBlogId(id uint64) ([]entity.BlogTag, error) {
	var blogtags []entity.BlogTag
	err := bt.db.Debug().
		Table("blogtags").
		Where("blog_id = ?", id).
		Find(&blogtags).
		Error
	if err != nil {
		return nil, err
	}
	return blogtags, err
}

func (bt *BlogTagRepo) GetByTagId(id uint64) ([]entity.BlogTag, error) {
	var blogtags []entity.BlogTag
	err := bt.db.Debug().
		Table("blogtags").
		Where("tag_id = ?", id).
		Find(&blogtags).
		Error
	if err != nil {
		return nil, err
	}
	return blogtags, err
}
