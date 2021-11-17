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

func (h *taskHandler) Index(c *gin.Context) {
	tasks, err := h.taskService.Index()
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
		Message: "Get all tasks",
		Data:    tasks,
	}
	c.JSON(http.StatusOK, response)
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
	c.JSON(http.StatusOK, response)
}

func (h *taskHandler) Show(c *gin.Context) {

	var input task.InputTaskDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	task, err := h.taskService.SelectById(input)
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
		Message: "Get task by id",
		Data:    task,
	}
	c.JSON(http.StatusOK, response)
}
