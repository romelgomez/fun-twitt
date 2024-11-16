package v1

import (
	controller "funtwitt/domain/follow/controller"

	"github.com/julienschmidt/httprouter"
)

func NewFollowRoutes(router *httprouter.Router, followController *controller.FollowController) {
	router.POST("/api/v1/follow", followController.Create)
	router.GET("/api/v1/follow/follower/:followerID", followController.FindByFollowerID)
	router.DELETE("/api/v1/follow/:id", followController.Delete)
}
