package modules

import (
	"famyun/pkg/web"
	"famyun/src/application/services"
	"famyun/src/infrastructure/dao"
	"famyun/src/interfaces/handlers"
	"go.uber.org/fx"
)

var Movie = fx.Module("MovieModel",
	fx.Provide(
		web.AsHandler(handlers.NewMovieHandler),
	),
	fx.Provide(
		dao.NewVideoRepo,
		dao.NewPerformerRepo,
		dao.NewMovieRepo,
	),
	fx.Provide(
		services.NewMovieService,
	),
)
