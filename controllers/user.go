package controllers

import (
	"net/http"

	m "echotest/models"
	"strconv"

	"github.com/labstack/echo"
)

// GetUsers gets all users
func GetUsers(c echo.Context) error {

	users := m.GetUsers()
	return c.JSON(http.StatusCreated, users)
}

// GetUser gets a specific user
func GetUser(c echo.Context) error {
	id := c.Param("id")

	i, _ := strconv.Atoi(id)
	users := m.GetUser(i)
	return c.JSON(http.StatusCreated, users)
}
