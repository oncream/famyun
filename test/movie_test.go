package test

import (
	"famyun/src/domain/models"
	"fmt"
	"testing"
)

func TestNewMovie(t *testing.T) {
	movies := models.NewVideo(
		models.WithVideoCnName("夏洛特烦恼"),
	)
	fmt.Println(movies)
}
