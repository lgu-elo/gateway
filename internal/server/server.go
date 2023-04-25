package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lsu-group/crm-api-gateway/docs"
	"github.com/sirupsen/logrus"
)

//	@title		Swagger Example API
//	@version	1.0
type CustomValidator struct {
	validator *validator.Validate
}

func New(logger *logrus.Logger) *echo.Echo {
	e := echo.New()

	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/healthz", func(c echo.Context) error { return c.String(http.StatusOK, "ok") })
	e.Use(newCors())
	e.Use(middleware.RequestLoggerWithConfig(setLogger(logger)))
	e.Use(setServerHeader)
	e.Pre(middleware.RemoveTrailingSlash())

	return e
}

func newCors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:*",
			"https://localhost:*",
			"http://127.0.0.1:*",
			"https://127.0.0.1:*",
			"http://1250753-co93798.tw1.ru",
			"https://1250753-co93798.tw1.ru",
			"http://*.1250753-co93798.tw1.ru",
			"https://*.1250753-co93798.tw1.ru",
		},
		Skipper: middleware.DefaultSkipper,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAccept,
			echo.HeaderCookie,
			echo.HeaderContentLength,
			echo.HeaderContentType,
			echo.HeaderCacheControl,
			echo.HeaderAcceptEncoding,
			echo.HeaderReferrerPolicy,
			echo.HeaderSetCookie,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowCredentials,
		},
		AllowMethods: []string{
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodOptions,
		},
		AllowCredentials: true,
	})
}

func setLogger(log *logrus.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func setServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "GATEWAY/1.0")
		return next(c)
	}
}
