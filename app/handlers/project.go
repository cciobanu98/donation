package handler

import (
	"net/http"

	"github.com/cciobanu98/donation/app/models/project"
	"github.com/gin-gonic/gin"
)

var handler *ProjectHandler

// ProjectHandler type
type ProjectHandler struct {
	projectRepo project.Repository
}

// CreateProjectInput struct
type CreateProjectInput struct {
	Title string
}

// NewHandler controller
func NewHandler(repo project.Repository) {
	handler = &ProjectHandler{repo}
}

// GetProjects godoc
// @Summary get projects
// @Schemes
// @Description get all projects
// @Tags project
// @Produce json
// @Success 200
// @Router /project [get]
func GetProjects(c *gin.Context) {
	projects, err := handler.projectRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": projects,
	})

}

// CreateProject godoc
// @Summary      create a new project
// @Description  create a new project
// @Tags project
// @Accept       json
// @Produce      json
// @Param        message  body      CreateProjectInput  true  "Project Info"
// @Success      200      
// @Failure      500      
// @Router       /project [post]
func CreateProject(c *gin.Context) {
	var input CreateProjectInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	project := project.Project{Title: input.Title}
	handler.projectRepo.Create(&project)
	
	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

// GetProject godoc
// @Summary Get project by Id
// @Description Get project using id
// @Tags project
// @Accept  json
// @Produce  json
// @Param   project_id     path    int     true        "Project ID"
// @Success 200 {string} string	"ok"
// @Failure 400
// @Failure 404
// @Router /project/{project_id} [get]
func GetProject(c *gin.Context) {
	id := c.Param("id")
	project, err := handler.projectRepo.Get(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": project,
	})
}

// Delete godoc
// @Summary delete project by Id
// @Description delete project using id
// @Tags project
// @Accept  json
// @Produce  json
// @Param   project_id     path    int     true        "Project ID"
// @Success 200 {string} string	"ok"
// @Failure 400
// @Failure 404
// @Router /project/{project_id} [delete]
func Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.projectRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})
}