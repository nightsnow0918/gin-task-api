package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"gin-task-api/database"
	"gin-task-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db *database.Database
}

func NewTaskHandler(db *database.Database) *TaskHandler {
	return &TaskHandler{db: db}
}

func SetupRoutes(r *gin.Engine, db *database.Database) *gin.Engine {
	taskHandler := NewTaskHandler(db)

	// Define routes
	r.GET("/tasks", taskHandler.GetAllTasks)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
	return r
}

//	@BasePath	/

// @Summary	Get all tasks
// @Schemes
// @Description	Get all tasks
// @Tags			Task
// @Accept			json
// @Produce		json
// @Success		200	{string}	GetAllTasks
// @Router			/tasks [get]
func (th *TaskHandler) GetAllTasks(c *gin.Context) {
	var tasks []models.Task
	result := th.db.DB.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

//	@BasePath	/

// @Summary	Create a task
// @Schemes
// @Description	Create a task with name and status
// @Tags			Task
// @Accept			json
// @Produce		json
// @param			name	body		string	true	"name"
// @param			status	body		int		true	"status"
// @Success		201		{string}	CreateTask
// @Router			/tasks [post]
func (th *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !(task.Status == models.COMPLETED || task.Status == models.INCOMPLETED) {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("incorrect value of parameter 'status'")})
		return
	}

	result := th.db.DB.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

//	@BasePath	/

// @Summary	Update task
// @Schemes
// @Description	Update a task by ID
// @Tags			Task
// @Accept			json
// @Produce		json
// @param			id		path		int		true	"id"
// @param			name	body		string	false	"name"
// @param			status	body		int		false	"status"
// @Success		200		{string}	UpdateTask
// @Router			/tasks/{id} [put]
func (th *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var existingTask models.Task
	result := th.db.DB.First(&existingTask, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Bind request body to update existingTask struct
	var updateTask models.Task
	if err := c.BindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update specific fields (consider using a patch method for partial updates)
	existingTask.Name = updateTask.Name
	existingTask.Status = updateTask.Status

	result = th.db.DB.Save(&existingTask)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, existingTask)
}

//	@BasePath	/

// @Summary	Delete task
// @Schemes
// @Description	Delete a task by ID
// @Tags			Task
// @Accept			json
// @Produce		json
// @param			id	path		int	true	"id"
// @Success		204	{string}	DeleteTask
// @Router			/tasks/{id} [delete]
func (th *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	result := th.db.DB.Delete(&models.Task{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
