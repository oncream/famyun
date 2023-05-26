package services

import (
	"famyun/pkg/logger"
	"famyun/src/application/assembler"
	"famyun/src/application/dto"
	"famyun/src/domain/repos"
)

type MovieService struct {
	movieRepo repos.IMovie
	log       logger.Logger
}

func NewMovieService(movieRepo repos.IMovie, log logger.New) *MovieService {
	return &MovieService{
		movieRepo: movieRepo,
		log:       log("famyun:MovieModule:MovieService"),
	}
}

func (m *MovieService) GetMovieByName(name string) (*dto.MovieDto, error) {
	movies, err := m.movieRepo.FindByName(name)
	if err != nil {
		m.log.Error(err)
		return nil, err
	}
	return assembler.MovieResp{}.ModelToDto(movies[0]), nil
}

func (m *MovieService) GetMovieById(id int64) (*dto.MovieDto, error) {
	movie, err := m.movieRepo.GetById(id)
	if err != nil {
		m.log.Error(err)
	}
	return assembler.MovieResp{}.ModelToDto(movie), nil
}

func (m *MovieService) FillTestData() error {
	err := m.movieRepo.FillTestData()
	if err != nil {
		m.log.Error(err)
	}
	return err
}
