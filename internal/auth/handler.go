package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lgu-elo/auth/pkg/jwt"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	e           *echo.Echo
	log         *logrus.Logger
	authService IService
}

func NewHandler(srv *echo.Echo, service IService, logger *logrus.Logger) *Handler {
	h := &Handler{
		srv,
		logger,
		service,
	}

	auth := srv.Group("auth")
	auth.POST("/login", h.SignIn)
	auth.POST("/register", h.SignUp)
	auth.POST("/this", h.This)

	return h
}

// Sign in service
//
//	@Summary	Login
//	@Tags		auth
//	@Accept		json
//	@Param		credentials	body	Credentials	true	"User creds to authenticate"
//	@Success	204			"User authenticated"
//	@Failure	400			"Bad request"
//	@Failure	500			"Internal server error"
//	@Router		/auth/login [post]
func (h *Handler) SignIn(c echo.Context) error {
	var creds Credentials

	if err := c.Bind(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := h.authService.Login(creds)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	originHeaders := c.Request().Header["Origin"]

	if len(originHeaders) != 0 {
		cookie := createJWTCookie(res, originHeaders[0])
		c.SetCookie(cookie)
	}

	return c.NoContent(http.StatusNoContent)
}

// Sign up service
//
//	@Summary	Register
//	@Tags		auth
//	@Accept		json
//	@Param		credentials	body	FullCredentials	true	"User creds to register"
//	@Success	204			"User created"
//	@Failure	400			"Bad request"
//	@Failure	500			"Internal server error"
//	@Router		/auth/register [post]
func (h *Handler) SignUp(c echo.Context) error {
	var creds FullCredentials

	if err := c.Bind(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	if err := c.Validate(&creds); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "")
	}

	res, err := h.authService.CreateNewUser(creds)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	originHeaders := c.Request().Header["Origin"]

	if len(originHeaders) != 0 {
		cookie := createJWTCookie(res, originHeaders[0])
		c.SetCookie(cookie)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) This(c echo.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	fmt.Print(token)

	id, err := jwt.ExtractData(token.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, id.ID)
}

func createJWTCookie(token, origin string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(7 * time.Hour * 24)
	cookie.HttpOnly = false
	cookie.Path = "/"
	cookie.Secure = false
	cookie.SameSite = http.SameSiteNoneMode

	if !strings.Contains(origin, "localhost") {
		cookie.SameSite = http.SameSiteLaxMode
	}

	return cookie
}
