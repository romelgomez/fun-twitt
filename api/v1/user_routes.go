package v1

import (
	controller "funtwitt/domain/user/controller"

	"github.com/julienschmidt/httprouter"
)

func NewUserRoutes(router *httprouter.Router, userController *controller.UserController) {
	router.POST("/api/v1/user", userController.Create)
	router.GET("/api/v1/user", userController.FindAll)
	router.GET("/api/v1/user/:id", userController.FindByID)
	router.PATCH("/api/v1/user", userController.Update)
	router.DELETE("/api/v1/user/:id", userController.Delete)
}
