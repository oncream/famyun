package repos

import (
	"famyun/src/domain/aggregates"
	"famyun/src/infrastructure/page"
)

type IMovie interface {
	Page(page *page.Page) error
	GetById(int64) (*aggregates.Movie, error)
	FindById(...int64) ([]*aggregates.Movie, error)
	FindByName(string) ([]*aggregates.Movie, error)
	Save(*aggregates.Movie) error
	Update(*aggregates.Movie) error
	FillTestData() error
}
