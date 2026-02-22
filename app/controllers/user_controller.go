package controllers

import (
	"errors"
	"net/http"
	"went-template/app/models"
	"went-template/internal/responses"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// GetAllUsers retrieves all users from the database.
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	var user models.User
	users, err := user.FindAll(uc.DB, map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by their ID.
// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} responses.ErrorResponse
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "ID parameter is required"})
		return
	}
	if err := user.FindByID(uc.DB, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user in the database.
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User to create"
// @Success 201 {object} models.User
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := user.Create(uc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser updates an existing user in the database.
// @Summary Update an existing user
// @Description Update an existing user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User to update"
// @Success 200 {object} models.User
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := user.FindByID(uc.DB, id); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: err.Error()})
		return
	}

	if err := user.Update(uc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user from the database.
// @Summary Delete a user
func (uc *UserController) DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Validate ID parameter
	if id == "" {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "ID parameter is required"})
		return
	}

	if err := user.FindByID(uc.DB, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		}
		return
	}

	if err := user.Delete(uc.DB); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
