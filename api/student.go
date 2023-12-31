package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentAPI interface {
	AddStudent(c *gin.Context)
	DeleteStudent(c *gin.Context)
}

type studentAPI struct {
	studentRepo repo.StudentRepository
}

func NewStudentAPI(studentRepo repo.StudentRepository) *studentAPI {
	return &studentAPI{studentRepo}
}

func (s *studentAPI) AddStudent(c *gin.Context) {
	var newStudent model.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := s.studentRepo.Store(&newStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add student success"})
}

func (s *studentAPI) DeleteStudent(c *gin.Context) {
	// TODO: answer here
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid ID"})
		return
	}

	err = s.studentRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Delete student success"})
}
