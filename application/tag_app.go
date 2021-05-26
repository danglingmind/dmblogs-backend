package application

import (
	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
)

type TagAppInterface interface {
	Save(tag entity.Tag) (*entity.Tag, error)
	DeleteById(id uint64) error
	DeleteByName(name string) error
	GetAllTags(limit, offset int) ([]entity.Tag, error)
	GetTagById(id uint64) (*entity.Tag, error)
	GetTagByName(name string) (*entity.Tag, error)
	GetTagsByIds(ids []uint64) ([]entity.Tag, error)
}

type TagApp struct {
	tr repository.TagRepository
}

var _ repository.TagRepository = &TagApp{}

func (t *TagApp) Save(tag entity.Tag) (*entity.Tag, error) {
	return t.tr.Save(tag)
}
func (t *TagApp) DeleteById(id uint64) error {
	return t.tr.DeleteById(id)
}
func (t *TagApp) DeleteByName(name string) error {
	return t.tr.DeleteByName(name)
}
func (t *TagApp) GetAllTags(limit, offset int) ([]entity.Tag, error) {
	return t.tr.GetAllTags(limit, offset)
}
func (t *TagApp) GetTagById(id uint64) (*entity.Tag, error) {
	return t.tr.GetTagById(id)
}
func (t *TagApp) GetTagByName(name string) (*entity.Tag, error) {
	return t.tr.GetTagByName(name)
}

func (t *TagApp) GetTagsByIds(ids []uint64) ([]entity.Tag, error) {
	return t.tr.GetTagsByIds(ids)
}
