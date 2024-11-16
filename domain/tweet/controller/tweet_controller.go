package tweet

import (
	dto "funtwitt/domain/tweet/dto"
	service "funtwitt/domain/tweet/service"
	"funtwitt/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TweetController struct {
	Service service.TweetService
}

func NewTweetController(service service.TweetService) *TweetController {
	return &TweetController{Service: service}
}

func (c *TweetController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data dto.TweetCreate
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

func (c *TweetController) FindByUserID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := params.ByName("id")
	result, err := c.Service.FindByUserID(r.Context(), userID)
	helper.Response(w, result, err)
}

func (c *TweetController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	err := c.Service.Delete(r.Context(), id)
	helper.Response(w, map[string]string{"message": "Tweet deleted successfully"}, err)
}
