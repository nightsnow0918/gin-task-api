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
	s.db.DB.Exec("DELETE FROM tasks")
}

func TestTaskAPISuite(t *testing.T) {
	suite.Run(t, new(TaskTestSuite))
}

func (s *TaskTestSuite) CreateTask(
	name string,
	status models.TaskStatus,
) models.Task {
	newTask := models.Task{Name: name, Status: status}
	s.db.DB.Create(&newTask)
	return newTask
}

func (s *TaskTestSuite) TestGetAllTasks() {
	var err error
	newTasks := []models.Task{
		{Name: "Task 1", Status: 0},
		{Name: "Task 2", Status: 1},
		{Name: "Task 3", Status: 1},
	}

	//Create a sample task
	for i := 0; i < len(newTasks); i++ {
		newTasks[i] = s.CreateTask(newTasks[i].Name, newTasks[i].Status)
	}

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
	s.Equal(len(newTasks), len(tasks))

	for i := 0; i < len(newTasks); i++ {
		s.Equal(newTasks[i].Id, tasks[i].Id)
		s.Equal(newTasks[i].Name, tasks[i].Name)
		s.Equal(newTasks[i].Status, tasks[i].Status)
	}
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
	createdTask := s.CreateTask("Task 1", 1)

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
	createdTask := s.CreateTask("Task 1", 1)

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
