package tweet

import (
	service "funtwitt/domain/timeline/service"
	"funtwitt/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TimelineController struct {
	Service service.TimelineService
}

func NewTimelineController(service service.TimelineService) *TimelineController {
	return &TimelineController{Service: service}
}

func (c *TimelineController) GetTimeline(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	result, err := c.Service.GetTimeline(r.Context(), id)
	helper.Response(w, result, err)
}
