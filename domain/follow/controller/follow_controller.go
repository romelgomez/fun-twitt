package follow

import (
	dto "funtwitt/domain/follow/dto"
	service "funtwitt/domain/follow/service"
	"funtwitt/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FollowController struct {
	Service service.FollowService
}

func NewFollowController(service service.FollowService) *FollowController {
	return &FollowController{Service: service}
}

func (c *FollowController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data dto.FollowCreate
	if err := helper.ReadRequestBody(r, &data); err != nil {
		helper.WriteErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := data.Validate(); err != nil {
		helper.WriteErrorResponse(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.Service.Create(r.Context(), data)
	helper.Response(w, result, err)
}

func (c *FollowController) FindByFollowerID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	followerID := params.ByName("followerID")
	result, err := c.Service.FindByFollowerID(r.Context(), followerID)
	helper.Response(w, result, err)
}

func (c *FollowController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	err := c.Service.Delete(r.Context(), id)
	helper.Response(w, map[string]string{"message": "Follow deleted successfully"}, err)
}
