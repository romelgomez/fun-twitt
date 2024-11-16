package router

import (
	"fmt"
	v1 "funtwitt/api/v1"

	user_controller "funtwitt/domain/user/controller"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userController *user_controller.UserController,
) *httprouter.Router {
	router := httprouter.New()

	// Home route
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	// Register v1 routes
	v1.NewUserRoutes(router, userController)

	// Register v2 routes (if needed)
	// v2.NewAccountRoutes(router, accountController)

	return router
}
