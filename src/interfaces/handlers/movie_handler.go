package handlers

import (
	"famyun/pkg/logger"
	"famyun/pkg/web"
	"famyun/pkg/web/errors"
	"famyun/src/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	movieService *services.MovieService
	log          logger.Logger
}

func NewMovieHandler(movieService *services.MovieService, log logger.New) *MovieHandler {
	return &MovieHandler{movieService: movieService, log: log("famyun:handlers:movieHandler")}
}

var _ web.Handler = (*MovieHandler)(nil)

func (m *MovieHandler) RegisterRouters(router gin.IRouter) {
	movie := router.Group("/movie")
	movie.GET("/get-by-name", m.GetMovieByName)
	movie.GET("/:id", m.GetMovieById)
	movie.GET("/testdata", m.FillTestData)
}

func (m *MovieHandler) GetMovieByName(c *gin.Context) {
	v, err := m.movieService.GetMovieByName(c.Query("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error500())
	} else {
		c.JSON(http.StatusOK, v)
	}
}

func (m *MovieHandler) GetMovieById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		m.log.Error(err)
		c.JSON(http.StatusInternalServerError, errors.Error500())
		return
	}
	v, err := m.movieService.GetMovieById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error500())
	} else {
		c.JSON(http.StatusOK, v)
	}
}

func (m *MovieHandler) FillTestData(c *gin.Context) {
	err := m.movieService.FillTestData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.Error500())
	} else {
		c.AbortWithStatus(http.StatusOK)
	}

}
