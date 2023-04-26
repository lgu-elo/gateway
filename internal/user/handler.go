package auth

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lgu-elo/user/pkg/pb"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	e           *echo.Echo
	log         *logrus.Logger
	userService IService
}

func NewHandler(srv *echo.Echo, service IService, logger *logrus.Logger) *Handler {
	h := &Handler{
		srv,
		logger,
		service,
	}

	users := srv.Group("users")
	users.GET("/:id", h.getUserById)
	users.DELETE("/:id", h.deleteUserById)
	users.PATCH("", h.updateUser)
	users.GET("", h.getAllUsers)
	users.POST("", h.createUser)

	return h
}

// Get user by id
//
//	@Summary	get by id
//	@Tags		user
//	@Accept		json
//	@Param		id	body		integer	true	"User id"
//	@Success	200	{object}	User
//	@Failure	404	"User not found"
//	@Failure	400	"Bad request"
//	@Router		/users/{id} [get]
func (h *Handler) getUserById(c echo.Context) error {

	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// Get all users
//
//	@Summary	get all
//	@Tags		user
//	@Success	200	{object}	[]User
//	@Failure	500	"Internal server error"
//	@Router		/users [get]
func (h *Handler) getAllUsers(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, users)
}

// Delete user by id
//
//	@Summary	delete by id
//	@Tags		user
//	@Accept		json
//	@Param		id	body	integer	true	"User id"
//	@Success	204	"User deleted"
//	@Failure	500	"Internal server error"
//	@Router		/users/{id} [delete]
func (h *Handler) deleteUserById(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.userService.DeleteUser(int64(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// Update user
//
//	@Summary	update user
//	@Tags		user
//	@Accept		json
//	@Param		user	body	User	true	"Updating user"
//	@Success	204		"User updated"
//	@Failure	500		"Internal server error"
//	@Router		/users [patch]
func (h *Handler) updateUser(c echo.Context) error {
	var user User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedUser, err := h.userService.UpdateUser(&pb.User{
		Id:    int64(user.ID),
		Login: user.Login,
		Name:  user.Name,
		Role:  user.Role,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// Create user
//
//	@Summary	create user
//	@Tags		user
//	@Accept		json
//	@Param		user	body	Creds	true	"create request"
//	@Success	201		"User created"
//	@Failure	500		"Internal server error"
//	@Failure	400		"Bad request"
//	@Router		/users [post]
func (h *Handler) createUser(c echo.Context) error {
	var user Creds

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userPb := &pb.User{
		Login:    user.Login,
		Password: user.Password,
		Name:     user.Name,
		Role:     user.Role,
	}

	err := h.userService.RegisterUser(userPb)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
