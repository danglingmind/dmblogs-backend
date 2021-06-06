package mock

import "danglingmind.com/ddd/domain/entity"

type BlogTagAppInterface struct {
	SaveFn           func(entity.BlogTag) error
	DeleteFn         func(uint64) error
	GetByBlogIdFn    func(uint64) ([]entity.BlogTag, error)
	GetByTagIdFn     func(uint64) ([]entity.BlogTag, error)
	DeleteByBlogIdFn func([]uint64) error
	DeleteByTagIdFn  func([]uint64) error
}

func (b *BlogTagAppInterface) Save(blogTag entity.BlogTag) error {
	return b.SaveFn(blogTag)
}

func (b *BlogTagAppInterface) Delete(id uint64) error {
	return b.DeleteFn(id)
}
func (b *BlogTagAppInterface) GetByBlogId(id uint64) ([]entity.BlogTag, error) {
	return b.GetByBlogIdFn(id)
}

func (b *BlogTagAppInterface) GetByTagId(id uint64) ([]entity.BlogTag, error) {
	return b.GetByTagIdFn(id)
}

func (b *BlogTagAppInterface) DeleteByBlogId(ids []uint64) error {
	return b.DeleteByBlogIdFn(ids)
}

func (b *BlogTagAppInterface) DeleteByTagId(ids []uint64) error {
	return b.DeleteByTagIdFn(ids)
}
