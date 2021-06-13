package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

type TaskType struct {
	Name string `json:"name"`
}

func CreateType (c echo.Context) error {
	taskType := new(TaskType)
	fmt.Println("creating", c)	
	if err := c.Bind(taskType); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, taskType)
}