package config

import (
	"context"
	"famyun/pkg/logger"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
	"net"
	"net/http"
	"strings"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Module = fx.Module(
	"base_providers",
	fx.Provide(
		NewAppConfig,
		NewGinEngine,
		NewLoggerFactory,
		NewXormEngine,
	),
)

var log = logger.NewLogger("famyun")

func NewGinEngine(lc fx.Lifecycle, config *AppConfig) *gin.Engine {
	if strings.ToLower(config.ServerConfig.Mode) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	handle := gin.New()
	handle.Use(gin.Recovery())
	srv := &http.Server{Addr: fmt.Sprintf(":%d", config.ServerConfig.Port), Handler: handle}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Infof("Web server starting on port %v", srv.Addr)
			log.Info("Press Ctrl + c to stop")
			go func() {
				if err0 := srv.Serve(ln); err0 != nil {
					log.Error(err0)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return handle
}

func NewXormEngine(config *AppConfig, l logger.New) *xorm.Engine {
	dbConfig := config.DatabaseConfig
	if dbConfig.Driver == "sqlite3" {
		engine, err := xorm.NewEngine("sqlite3", fmt.Sprintf("./%s.db", dbConfig.Database))
		if err != nil {
			log.Fatal(err)
		}
		engine.SetTableMapper(names.NewPrefixMapper(names.GonicMapper{}, "pm_"))
		engine.SetColumnMapper(names.GonicMapper{})
		engine.SetLogger(logger.NewXormLogger(l("xorm")))
		if dbConfig.ShowSql {
			engine.ShowSQL(true)
		}
		if err = backoff.Retry(engine.Ping, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10)); err != nil {
			log.Fatal(err)
		}
		return engine
	} else {
		url := fmt.Sprintf("%s:%s@(%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Database)
		engine, err := xorm.NewEngineWithParams("mysql", url, dbConfig.Params)
		if err != nil {
			log.Fatal(err)
		}
		engine.SetTableMapper(names.NewPrefixMapper(names.GonicMapper{}, "pm_"))
		engine.SetColumnMapper(names.GonicMapper{})
		engine.SetLogger(logger.NewXormLogger(l("xorm")))
		if dbConfig.ShowSql {
			engine.ShowSQL(true)
		}
		if err = backoff.Retry(engine.Ping, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10)); err != nil {
			log.Fatal(err)
		}
		return engine
	}
}

func NewLoggerFactory(config *AppConfig) logger.New {
	return func(name string) logger.Logger {
		return logger.NewLogger(name)
	}
}
