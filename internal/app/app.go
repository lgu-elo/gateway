package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	auth "github.com/lgu-elo/gateway/internal/auth"
	"github.com/lgu-elo/gateway/internal/config"
	"github.com/lgu-elo/gateway/internal/server"
	user "github.com/lgu-elo/gateway/internal/user"
	"github.com/lgu-elo/gateway/pkg/logger"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	fxlogrus "github.com/takt-corp/fx-logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Run() {
	fx.New(CreateApp()).Run()
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.WithLogger(func(log *logrus.Logger) fxevent.Logger {
			return &fxlogrus.LogrusLogger{Logger: log}
		}),
		fx.Provide(
			logger.New,
			config.New,

			server.New,

			fx.Annotate(user.NewClient, fx.As(new(user.Client))),
			fx.Annotate(user.NewService, fx.As(new(user.IService))),

			fx.Annotate(auth.NewClient, fx.As(new(auth.Client))),
			fx.Annotate(auth.NewService, fx.As(new(auth.IService))),

			//fx.Annotate(project.NewClient, fx.As(new(project.Client))),
			//fx.Annotate(project.NewService, fx.As(new(project.IService))),

			//fx.Annotate(task.NewClient, fx.As(new(task.Client))),
			//fx.Annotate(task.NewService, fx.As(new(task.IService))),

			auth.NewHandler,
			user.NewHandler,
			//project.NewHandler,
			//task.NewHandler,
		),
		fx.Invoke(serve),
	)
}

func serve(e *echo.Echo, cfg *config.Cfg, ah *auth.Handler /*, uh *user.Handler, ih *project.Handler, dh *task.Handler */) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(
		":%s",
		cfg.Port,
	)))
}
