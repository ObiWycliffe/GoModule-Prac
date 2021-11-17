package handler

import (
	"net/http"

	"github.com/ObiWycliffe/GoModule/go_crud/task"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService}
}

func (handler *taskHandler) Store(c *gin.Context) {
	var input task.InputTask
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTask, err := handler.taskService.Store(input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Message: "New task has been stored successfully",
		Data:    newTask,
	}
	c.JSON(http.StatusBadRequest, response)
}
