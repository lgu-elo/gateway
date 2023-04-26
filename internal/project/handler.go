// package project

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// 	"github.com/lgu-elo/project/pkg/pb"
// 	"github.com/sirupsen/logrus"
// )

// type Handler struct {
// 	e              *echo.Echo
// 	log            *logrus.Logger
// 	projectService IService
// }

// func NewHandler(srv *echo.Echo, service IService, logger *logrus.Logger) *Handler {
// 	h := &Handler{
// 		srv,
// 		logger,
// 		service,
// 	}

// 	projects := srv.Group("projects")
// 	projects.GET("/:id", h.getProjectById)
// 	projects.DELETE("/:id", h.deleteProjectById)
// 	projects.PATCH("", h.updateProject)
// 	projects.GET("", h.getAllProjects)
// 	projects.POST("", h.createProject)

// 	return h
// }

// // Get project by id
// //
// //	@Summary	get by id
// //	@Tags		project
// //	@Accept		json
// //	@Param		id	body		integer	true	"Project id"
// //	@Success	200	{object}	Project
// //	@Failure	404	"Project not found"
// //	@Failure	400	"Bad request"
// //	@Router		/projects/{id} [get]
// func (h *Handler) getProjectById(c echo.Context) error {

// 	id := c.Param("id")
// 	projectID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	project, err := h.projectService.GetProjectById(projectID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, project)
// }

// // Get all projects
// //
// //	@Summary	get all
// //	@Tags		project
// //	@Success	200	{object}	[]Project
// //	@Failure	500	"Internal server error"
// //	@Router		/projects [get]
// func (h *Handler) getAllProjects(c echo.Context) error {
// 	projects, err := h.projectService.GetAllProjects()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "")
// 	}

// 	return c.JSON(http.StatusOK, projects)
// }

// // Delete project by id
// //
// //	@Summary	delete by id
// //	@Tags		project
// //	@Accept		json
// //	@Param		id	body	integer	true	"Project id"
// //	@Success	204	"Project deleted"
// //	@Failure	500	"Internal server error"
// //	@Router		/projects/{id} [delete]
// func (h *Handler) deleteProjectById(c echo.Context) error {
// 	id := c.Param("id")
// 	projectID, err := strconv.Atoi(id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	err = h.projectService.DeleteProject(int64(projectID))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, err.Error())
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }

// // Update project
// //
// //	@Summary	update project
// //	@Tags		project
// //	@Accept		json
// //	@Param		project	body	Project	true	"Updating user"
// //	@Success	204		"Project updated"
// //	@Failure	500		"Internal server error"
// //	@Router		/projects [patch]
// func (h *Handler) updateProject(c echo.Context) error {
// 	var project Project

// 	if err := c.Bind(&project); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := c.Validate(&project); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	updatedProject, err := h.projectService.UpdateProject(&pb.Project{
// 		Name:        project.Name,
// 		Description: project.Description,
// 	})
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "")
// 	}

// 	return c.JSON(http.StatusOK, updatedProject)
// }

// // Create project
// //
// //	@Summary	create project
// //	@Tags		project
// //	@Accept		json
// //	@Param		project	body	ProjectDto	true	"create request"
// //	@Success	201		"Project created"
// //	@Failure	500		"Internal server error"
// //	@Failure	400		"Bad request"
// //	@Router		/projects [post]
// func (h *Handler) createProject(c echo.Context) error {
// 	var project ProjectDto

// 	if err := c.Bind(&project); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := c.Validate(&project); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	projectPb := &pb.Project{
// 		Name:        project.Name,
// 		Description: project.Description,
// 	}

// 	err := h.projectService.CreateProject(projectPb)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.NoContent(http.StatusCreated)
// }
