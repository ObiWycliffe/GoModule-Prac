package handler

import (
	// "github.com/Obiycliffe/GoModule/go_crud/task"
	"../task"
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

func (handler *taskHandler) Store(c *gin.Context) {}
