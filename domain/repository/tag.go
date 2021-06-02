package repository

import "danglingmind.com/ddd/domain/entity"

type TagRepository interface {
	Save(tag entity.Tag) (*entity.Tag, error)
	DeleteById(id uint64) error
	DeleteByName(name string) error
	GetAllTags(limit, offset int) ([]entity.Tag, error)
	GetTagById(id uint64) (*entity.Tag, error)
	GetTagByName(name string) (*entity.Tag, error)
	GetTagsByIds(ids []uint64) ([]entity.Tag, error)
}
