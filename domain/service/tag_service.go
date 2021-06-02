package service

import (
	"fmt"

	"danglingmind.com/ddd/domain/entity"
	"danglingmind.com/ddd/domain/repository"
)

type TagServiceInterface interface {
	GetTagsByBlogId(id uint64) ([]entity.Tag, error)
}

type TagService struct {
	tg  repository.TagRepository
	btg repository.BlogTagRepository
}

var _ TagServiceInterface = &TagService{}

func NewTagService(t repository.TagRepository, btg repository.BlogTagRepository) *TagService {
	return &TagService{
		tg:  t,
		btg: btg,
	}
}

func (t *TagService) GetTagsByBlogId(id uint64) ([]entity.Tag, error) {
	// get the tag ids
	tagMappings, err := t.btg.GetByBlogId(id)
	if err != nil {
		return nil, fmt.Errorf("error while finding tags")
	}
	ids := make([]uint64, 0)
	for _, i := range tagMappings {
		ids = append(ids, i.TagId)
	}
	tags, err := t.tg.GetTagsByIds(ids)
	if err != nil {
		return nil, fmt.Errorf("error in finding tags")
	}
	return tags, nil
}
