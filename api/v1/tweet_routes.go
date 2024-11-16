package v1

import (
	controller "funtwitt/domain/tweet/controller"

	"github.com/julienschmidt/httprouter"
)

func NewTweetRoutes(router *httprouter.Router, tweetController *controller.TweetController) {
	router.POST("/api/v1/tweet", tweetController.Create)
	router.GET("/api/v1/tweet/user/:id", tweetController.FindByUserID)
	router.DELETE("/api/v1/tweet/:id", tweetController.Delete)
}
