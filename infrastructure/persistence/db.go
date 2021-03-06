package persistence

import (
	"fmt"
	"os"

	"danglingmind.com/ddd/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	User    repository.UserRepository
	Blog    repository.BlogRepository
	Tag     repository.TagRepository
	BlogTag repository.BlogTagRepository
	db      *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := os.Getenv("DATABASE_URL")
	if DBURL == "" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	}

	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return &Repositories{
		User:    NewUserRepository(db),
		Blog:    NewBlogRepository(db),
		Tag:     NewTagRepo(db),
		BlogTag: NewBlogTagRepo(db),
		db:      db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

//This migrate all tables
// func (s *Repositories) Automigrate() error {
// 	return s.db.AutoMigrate(&entity.User{}, &entity.Blog{}).Error
// }
