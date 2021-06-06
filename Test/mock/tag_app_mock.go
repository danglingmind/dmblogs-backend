package mock

import "danglingmind.com/ddd/domain/entity"

type TagAppInterface struct {
	SaveFn         func(entity.Tag) (*entity.Tag, error)
	DeleteByIdFn   func(uint64) error
	DeleteByNameFn func(string) error
	GetAllTagsFn   func(int, int) ([]entity.Tag, error)
	GetTagByIdFn   func(uint64) (*entity.Tag, error)
	GetTagByNameFn func(string) (*entity.Tag, error)
	GetTagsByIdsFn func([]uint64) ([]entity.Tag, error)
}

func (ta *TagAppInterface) Save(tag entity.Tag) (*entity.Tag, error) {
	return ta.SaveFn(tag)
}

func (ta *TagAppInterface) DeleteById(id uint64) error {
	return ta.DeleteByIdFn(id)
}
func (ta *TagAppInterface) DeleteByName(name string) error {
	return ta.DeleteByNameFn(name)
}
func (ta *TagAppInterface) GetAllTags(limit int, offset int) ([]entity.Tag, error) {
	return ta.GetAllTagsFn(limit, offset)
}
func (ta *TagAppInterface) GetTagById(id uint64) (*entity.Tag, error) {
	return ta.GetTagByIdFn(id)
}
func (ta *TagAppInterface) GetTagByName(name string) (*entity.Tag, error) {
	return ta.GetTagByNameFn(name)
}
func (ta *TagAppInterface) GetTagsByIds(ids []uint64) ([]entity.Tag, error) {
	return ta.GetTagsByIdsFn(ids)
}
