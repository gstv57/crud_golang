package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gstv57/crud_golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/get-user-by-id/:user_id", controller.FindUserById)
	r.GET("/get-user-by-email/:user_email", controller.FindUserByEmail)
	r.POST("/create/user", controller.CreateUser)
	r.PUT("/update-user/:user_id", controller.UpdateUser)
	r.DELETE("/delete-user/:user_id", controller.DeleteUser)
}
