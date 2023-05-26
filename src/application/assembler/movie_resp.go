package assembler

import (
	"famyun/src/application/dto"
	"famyun/src/domain/aggregates"
)

type MovieResp struct {
}

func NewMovieResp() *MovieResp {
	return &MovieResp{}
}

//todo 改成反射映射，使用json，不然models层的json标记无用。

func (this MovieResp) ModelToDto(movie *aggregates.Movie) *dto.MovieDto {
	if movie == nil {
		return nil
	}
	performer := make([]string, 0, len(movie.Performer))
	for _, p := range movie.Performer {
		performer = append(performer, p.Name)
	}
	return &dto.MovieDto{
		Name:        movie.Video.CnName,
		Synopsis:    movie.Video.Synopsis,
		ReleaseTime: movie.Video.ReleaseTime.Format("2006-01-02"),
		Performer:   performer,
		Director:    movie.Director.Name,
		Cover:       movie.Video.Cover,
		Address:     movie.Video.Address,
	}
}
