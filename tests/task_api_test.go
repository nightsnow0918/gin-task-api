package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-task-api/database"
	"gin-task-api/handlers"
	"gin-task-api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TaskTestSuite struct {
	suite.Suite

	router *gin.Engine
	db     *database.Database
}

func (s *TaskTestSuite) SetupTest() {
	// Setup in-memory database and router
	var err error

	s.db, err = database.InitDB()
	if err != nil {
		s.T().Fatal(err)
	}
	s.router = handlers.SetupRoutes(gin.Default(), s.db)
}

func (s *TaskTestSuite) TearDownTest() {
	// Clean-up table after each test
	s.db.DB.Where("1 = 1").Delete(&models.Task{})
}

func TestTaskAPISuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}

func (s *TaskTestSuite) TestGetAllTasks() {
	var err error

	//Create a sample task
	task1 := models.Task{Name: "Task 1", Status: 0}
	s.db.DB.Create(&task1)

	// Test start
	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		s.T().Fatal(err)
	}
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	var tasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &tasks)
	if err != nil {
		s.T().Fatal(err)
	}

	// Assertions
	s.Equal(http.StatusOK, w.Code)
	s.Equal(1, len(tasks))
	s.Equal(task1.Id, tasks[0].Id)
	s.Equal(task1.Name, tasks[0].Name)
	s.Equal(task1.Status, tasks[0].Status)
}

func (s *TaskTestSuite) TestCreateTasks() {
	var err error

	newTask := models.Task{
		Name:   "Test Task",
		Status: 1,
	}
	jsonData, err := json.Marshal(newTask)
	if err != nil {
		s.T().Fatal(err)
	}

	// Test start
	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(jsonData))
	if err != nil {
		s.T().Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	var createdTask models.Task
	err = json.Unmarshal(w.Body.Bytes(), &createdTask)
	if err != nil {
		s.T().Fatal(err)
	}

	// Assertions
	s.Equal(http.StatusCreated, w.Code)
	s.Equal(1, createdTask.Id)
	s.Equal(newTask.Name, createdTask.Name)
	s.Equal(newTask.Status, createdTask.Status)
}

func (s *TaskTestSuite) TestUpdateTasks() {
	var err error

	// Create a sample task
	createdTask := models.Task{Name: "Task 1", Status: 1}
	s.db.DB.Create(&createdTask)

	newTask := models.Task{
		Name:   "Task 1",
		Status: 0,
	}
	jsonData, err := json.Marshal(newTask)
	if err != nil {
		s.T().Fatal(err)
	}

	// Test start
	url := fmt.Sprintf("/tasks/%v", createdTask.Id)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(jsonData))
	if err != nil {
		s.T().Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	var updatedTask models.Task
	err = json.Unmarshal(w.Body.Bytes(), &updatedTask)
	if err != nil {
		s.T().Fatal(err)
	}

	// Assertions
	s.Equal(http.StatusOK, w.Code)
	s.Equal(newTask.Name, updatedTask.Name)
	s.Equal(newTask.Status, updatedTask.Status)
}

func (s *TaskTestSuite) TestDeleteTasks() {
	var err error

	// Create a sample task
	createdTask := models.Task{Name: "Task 1", Status: 1}
	s.db.DB.Create(&createdTask)

	url := fmt.Sprintf("/tasks/%v", createdTask.Id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		s.T().Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	// Assertions
	s.Equal(http.StatusNoContent, w.Code)
}
