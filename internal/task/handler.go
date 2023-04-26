package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lgu-elo/task/pkg/pb"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	e           *echo.Echo
	log         *logrus.Logger
	taskService IService
}

func NewHandler(srv *echo.Echo, service IService, logger *logrus.Logger) *Handler {
	h := &Handler{
		srv,
		logger,
		service,
	}

	tasks := srv.Group("tasks")
	tasks.GET("/:id", h.getTaskById)
	tasks.DELETE("/:id", h.deleteTaskById)
	tasks.PATCH("", h.updateTask)
	tasks.GET("", h.getAllTasks)
	tasks.POST("", h.createTask)

	return h
}

// Get task by id
//
//	@Summary	get by id
//	@Tags		task
//	@Accept		json
//	@Param		id	body		integer	true	"Task id"
//	@Success	200	{object}	Task
//	@Failure	404	"Task not found"
//	@Failure	400	"Bad request"
//	@Router		/tasks/{id} [get]
func (h *Handler) getTaskById(c echo.Context) error {

	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskService.GetTaskById(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

// Get all tasks
//
//	@Summary	get all
//	@Tags		task
//	@Success	200	{object}	[]Task
//	@Failure	500	"Internal server error"
//	@Router		/tasks [get]
func (h *Handler) getAllTasks(c echo.Context) error {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, tasks)
}

// Delete task by id
//
//	@Summary	delete by id
//	@Tags		task
//	@Accept		json
//	@Param		id	body	integer	true	"Task id"
//	@Success	204	"Task deleted"
//	@Failure	500	"Internal server error"
//	@Router		/tasks/{id} [delete]
func (h *Handler) deleteTaskById(c echo.Context) error {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.taskService.DeleteTask(int64(taskID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// Update task
//
//	@Summary	update task
//	@Tags		task
//	@Accept		json
//	@Param		task	body	Task	true	"Updating task"
//	@Success	204		"Task updated"
//	@Failure	500		"Internal server error"
//	@Router		/tasks [patch]
func (h *Handler) updateTask(c echo.Context) error {
	var task Task

	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updatedTask, err := h.taskService.UpdateTask(&pb.Task{
		Id:          int64(task.ID),
		Name:        task.Name,
		Description: task.Description,
		UserId:      int64(task.User_id),
		ProjectId:   int64(task.Project_id),
		Status:      task.Status,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedTask)
}

// Create task
//
//	@Summary	create task
//	@Tags		task
//	@Accept		json
//	@Param		task	body	TaskDto	true	"create task"
//	@Success	201		"Task created"
//	@Failure	500		"Internal server error"
//	@Failure	400		"Bad request"
//	@Router		/tasks [post]
func (h *Handler) createTask(c echo.Context) error {
	var task TaskDto

	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	pbTask := &pb.Task{
		Name:        task.Name,
		Description: task.Description,
		Status:      task.Status,
		UserId:      int64(task.User_id),
		ProjectId:   int64(task.Project_id),
	}

	err := h.taskService.CreateTask(pbTask)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
