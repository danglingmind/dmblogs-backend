package mock

import "danglingmind.com/ddd/domain/entity"

type TagServiceInterface struct {
	GetTagsByBlogIdFn func(uint64) ([]entity.Tag, error)
}

func (t *TagServiceInterface) GetTagsByBlogId(id uint64) ([]entity.Tag, error) {
	return t.GetTagsByBlogIdFn(id)
}
