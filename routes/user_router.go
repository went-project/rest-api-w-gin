package routes

import (
	"went-template/app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {

	uc := &controllers.UserController{DB: db}

	r.GET("/users", uc.GetAllUsers)
	r.GET("/users/:id", uc.GetUserByID)
	r.POST("/users", uc.CreateUser)
	r.PUT("/users/:id", uc.UpdateUser)
	r.DELETE("/users/:id", uc.DeleteUser)
}
