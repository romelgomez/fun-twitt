package v1

import (
	controller "funtwitt/domain/timeline/controller"

	"github.com/julienschmidt/httprouter"
)

func NewTimelineRoutes(router *httprouter.Router, timelineController *controller.TimelineController) {
	router.GET("/api/v1/timeline/:id", timelineController.GetTimeline)
}
