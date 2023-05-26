package modules

import (
	"famyun/pkg/logger"
	"famyun/src/domain/models"
	"go.uber.org/fx"
	"xorm.io/xorm"
)

var Module = fx.Module("module",
	fx.Options(
		Movie,
	),
	fx.Invoke(
		registerModels,
	),
)

func registerModels(engine *xorm.Engine, logger logger.New) {
	log := logger("famyun")
	if err := engine.Sync2(
		&models.Video{},
		&models.Performer{},
	); err != nil {
		log.Error(err)
	}
}
