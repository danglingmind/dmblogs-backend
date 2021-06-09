package persistence

import (
	"strings"

	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
	"github.com/jinzhu/gorm"
)

type TagRepo struct {
	db *gorm.DB
}

var _ repository.TagRepository = &TagRepo{}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{
		db: db,
	}
}
func (t *TagRepo) Save(tag entity.Tag) (*entity.Tag, error) {
	err := tag.PreSave()
	if err != nil {
		return nil, err
	}
	err = t.db.Debug().Table("tags").Create(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, err
}

func (t *TagRepo) DeleteById(id uint64) error {
	var tag entity.Tag
	tag.ID = id
	err := t.db.Debug().Table("tags").Delete(&tag).Error
	if err != nil {
		return err
	}
	return err
}

func (t *TagRepo) DeleteByName(name string) error {
	err := t.db.Debug().
		Table("tags").
		Delete(entity.Tag{}, "name like lower(?)", strings.ToLower(name)).Error
	if err != nil {
		return err
	}
	return err
}

func (t *TagRepo) GetAllTags(limit, offset int) ([]entity.Tag, error) {
	var tags []entity.Tag
	err := t.db.Debug().
		Table("tags").
		Limit(limit).
		Offset(offset).
		Error
	if err != nil {
		return nil, err
	}
	return tags, err
}

func (t *TagRepo) GetTagById(id uint64) (*entity.Tag, error) {
	var tag entity.Tag
	err := t.db.Debug().
		Table("tags").
		Where("id = ?", id).
		Find(&tag).
		Error
	return &tag, err
}

func (t *TagRepo) GetTagByName(name string) (*entity.Tag, error) {
	var tag entity.Tag
	err := t.db.Debug().
		Table("tags").
		Where("name like lower(?)", strings.ToLower(name)).
		Find(&tag).
		Error
	if err != nil {
		return nil, err
	}
	return &tag, err
}

func (t *TagRepo) GetTagsByIds(ids []uint64) ([]entity.Tag, error) {
	// gorm does not support the slice we have to prepare the IN statement
	inStatement := "id IN ("
	var params []interface{}
	for idx, i := range ids {
		params = append(params, i)
		if idx != len(ids)-1 {
			inStatement += "?,"
		} else {
			inStatement += "?"
		}
	}
	inStatement += ")"
	var tags []entity.Tag
	err := t.db.Debug().
		Table("tags").
		Where(inStatement, params...).
		Find(&tags).
		Error
	return tags, err
}
