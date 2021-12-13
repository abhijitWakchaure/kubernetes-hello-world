package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerListTasks(c *gin.Context) {
	tasks, err := listTasks()
	if err != nil {
		errorResponse := ResponseError{
			Code:  http.StatusInternalServerError,
			Error: "Failed to list tasks due to: " + err.Error(),
		}
		c.JSON(errorResponse.Code, errorResponse)
		return
	}
	response := ResponseListTasks{
		Meta:  meta,
		Tasks: tasks,
	}
	c.JSON(http.StatusOK, response)
}

func handlerInsertTask(c *gin.Context) {
	var request RequestInsertTask
	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorResponse := ResponseError{
			Code:  http.StatusBadRequest,
			Error: "Failed to bind JSON due to: " + err.Error(),
		}
		c.JSON(errorResponse.Code, errorResponse)
		return
	}
	task, err := insertTask(Task{Name: request.Name})
	if err != nil {
		errorResponse := ResponseError{
			Code:  http.StatusInternalServerError,
			Error: "Failed to insert task due to: " + err.Error(),
		}
		c.JSON(errorResponse.Code, errorResponse)
		return
	}
	c.JSON(http.StatusCreated, task)
}
