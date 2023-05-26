package dao

import (
	"famyun/src/domain/aggregates"
	"famyun/src/domain/models"
	"famyun/src/domain/repos"
	"famyun/src/infrastructure/page"
	"fmt"
	"xorm.io/xorm"
)

type MovieRepo struct {
	engine        *xorm.Engine
	performerRepo repos.IPerformer
	videoRepo     repos.IVideo
}

var _ repos.IMovie = (*MovieRepo)(nil)

func NewMovieRepo(engine *xorm.Engine, performerRepo repos.IPerformer, videoRepo repos.IVideo) repos.IMovie {
	return &MovieRepo{engine, performerRepo, videoRepo}
}

func (m MovieRepo) Page(page *page.Page) error {
	return nil
}

func (m MovieRepo) GetById(i int64) (*aggregates.Movie, error) {
	v, err := m.videoRepo.GetById(i)
	if err != nil {
		return nil, err
	}
	director, err := m.performerRepo.GetById(v.DirectorId)
	if err != nil {
		return nil, err
	}
	performer, err := m.performerRepo.FindById(v.PerformerIds...)
	if err != nil {
		return nil, err
	}
	return aggregates.NewMovie(
		aggregates.WithMovieVideo(v),
		aggregates.WithMovieDirector(director),
		aggregates.WithMoviePerformer(performer...),
	), nil

}

func (m MovieRepo) FindById(i ...int64) ([]*aggregates.Movie, error) {
	return nil, nil
}

func (m MovieRepo) FindByName(s string) ([]*aggregates.Movie, error) {
	return nil, nil
}

func (m MovieRepo) Save(movie *aggregates.Movie) error {
	return nil
}

func (m MovieRepo) Update(movie *aggregates.Movie) error {
	return nil
}

func (m MovieRepo) FillTestData() error {
	performer := []*models.Performer{
		models.NewPerformer(
			models.WithPerformerName("弗兰克·德拉邦特"),
			models.WithPerformerGender(models.Male),
			models.WithPerformerPhoto("https://img1.doubanio.com/view/celebrity/raw/public/p230.jpg"),
		),
		models.NewPerformer(
			models.WithPerformerName("蒂姆·罗宾斯"),
			models.WithPerformerGender(models.Male),
			models.WithPerformerPhoto("https://img9.doubanio.com/view/celebrity/raw/public/p17525.jpg"),
		),
		models.NewPerformer(
			models.WithPerformerName("摩根·弗里曼"),
			models.WithPerformerGender(models.Male),
			models.WithPerformerPhoto("https://img2.doubanio.com/view/celebrity/raw/public/p34642.jpg"),
		),
	}
	err := m.performerRepo.Save(performer...)
	if err != nil {
		return err
	}
	video := models.NewVideo(
		models.WithVideoCnName("肖申克的救赎"),
		models.WithVideoEnName("The Shawshank Redemption"),
		models.WithVideoCover("https://img2.doubanio.com/view/photo/s_ratio_poster/public/p480747492.webp"),
		models.WithVideoDirectorId(performer[0].Id),
		models.WithVideoPerformerIds(performer[1].Id, performer[2].Id),
	)
	fmt.Println(video)
	err = m.videoRepo.Save(video)
	if err != nil {
		return err
	}
	return err
}
