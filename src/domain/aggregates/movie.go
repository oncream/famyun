package aggregates

import (
	"famyun/src/domain/models"
	"time"
)

//电影聚合 视频+导演+演员

type Movie struct {
	Video     *models.Video //聚合根
	Director  *models.Performer
	Performer []*models.Performer
}

func NewMovie(cfg ...movieConfiguration) *Movie {
	video := models.NewVideo()
	movie := &Movie{Video: video}
	movieConfigurations(cfg).apply(movie)
	return movie
}

type movieConfiguration func(m *Movie)

type movieConfigurations []movieConfiguration

func (cfg movieConfigurations) apply(m *Movie) {
	for _, c := range cfg {
		c(m)
	}
}

func WithMovieCnName(arg string) movieConfiguration {
	return func(m *Movie) {
		m.Video.CnName = arg
	}
}

func WithMovieEnName(arg string) movieConfiguration {
	return func(m *Movie) {
		m.Video.EnName = arg
	}
}

func WithMovieSynopsis(arg string) movieConfiguration {
	return func(m *Movie) {
		m.Video.Synopsis = arg
	}
}

func WithMovieCover(arg string) movieConfiguration {
	return func(m *Movie) {
		m.Video.Cover = arg
	}
}

func WithMovieReleaseTime(arg time.Time) movieConfiguration {
	return func(m *Movie) {
		m.Video.ReleaseTime = arg
	}
}

func WithMovieAddress(arg string) movieConfiguration {
	return func(m *Movie) {
		m.Video.Address = arg
	}
}

func WithMovieVideo(arg *models.Video) movieConfiguration {
	return func(m *Movie) {
		m.Video = arg
	}
}

func WithMovieDirector(arg *models.Performer) movieConfiguration {
	return func(m *Movie) {
		m.Director = arg
	}
}

func WithMoviePerformer(arg ...*models.Performer) movieConfiguration {
	return func(m *Movie) {
		m.Performer = arg
	}
}
