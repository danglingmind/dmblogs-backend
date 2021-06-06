package mock

import (
	"danglingmind.com/ddd/domain/entity"
	"github.com/google/uuid"
)

type BlogAppInterface struct {
	SaveFn             func(entity.Blog, uuid.UUID) (*entity.Blog, error)
	DeleteFn           func(uint64) error
	GetBlogByIdFn      func(uint64) (*entity.Blog, error)
	GetBlogsFn         func(int, int) ([]entity.Blog, error)
	GetBlogsByIdsFn    func([]uint64, int, int) ([]entity.Blog, error)
	GetBlogsByUserIdFn func(uuid.UUID) ([]entity.Blog, error)
}

func (b *BlogAppInterface) Save(blog entity.Blog, userid uuid.UUID) (*entity.Blog, error) {
	return b.SaveFn(blog, userid)
}
func (b *BlogAppInterface) Delete(id uint64) error {
	return b.DeleteFn(id)
}
func (b *BlogAppInterface) GetBlogById(id uint64) (*entity.Blog, error) {
	return b.GetBlogByIdFn(id)
}
func (b *BlogAppInterface) GetBlogs(limit, offset int) ([]entity.Blog, error) {
	return b.GetBlogsFn(limit, offset)
}
func (b *BlogAppInterface) GetBlogsByIds(blogIds []uint64, limit, offset int) ([]entity.Blog, error) {
	return b.GetBlogsByIdsFn(blogIds, limit, offset)
}
func (b *BlogAppInterface) GetBlogsByUserId(userid uuid.UUID) ([]entity.Blog, error) {
	return b.GetBlogsByUserIdFn(userid)
}
