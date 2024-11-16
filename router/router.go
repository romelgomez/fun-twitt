package router

import (
	"fmt"
	v1 "funtwitt/api/v1"

	follow_controller "funtwitt/domain/follow/controller"
	timeline_controller "funtwitt/domain/timeline/controller"
	tweet_controller "funtwitt/domain/tweet/controller"
	user_controller "funtwitt/domain/user/controller"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userController *user_controller.UserController,
	tweetController *tweet_controller.TweetController,
	followController *follow_controller.FollowController,
	timelineController *timeline_controller.TimelineController,
) *httprouter.Router {
	router := httprouter.New()

	// Home route
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	// Register v1 routes
	v1.NewUserRoutes(router, userController)
	v1.NewTweetRoutes(router, tweetController)
	v1.NewFollowRoutes(router, followController)
	v1.NewTimelineRoutes(router, timelineController)

	// Register v2 routes (if needed)
	// v2.NewAccountRoutes(router, accountController)

	return router
}
