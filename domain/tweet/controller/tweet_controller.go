package tweet

import (
	"context"
	dto "funtwitt/domain/tweet/dto"
	service "funtwitt/domain/tweet/service"
	user_dto "funtwitt/domain/user/dto"
	user_service "funtwitt/domain/user/service"
	"funtwitt/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TweetController struct {
	Service     service.TweetService
	UserService user_service.UserService
}

func NewTweetController(service service.TweetService, userService user_service.UserService) *TweetController {
	return &TweetController{
		Service:     service,
		UserService: userService,
	}
}

func (c *TweetController) findUser(ctx context.Context, id string) (user_dto.UserResponse, error) {
	return c.UserService.FindByID(ctx, id)
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

	user, err := c.findUser(r.Context(), data.UserID)
	if err != nil {
		helper.WriteErrorResponse(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	data.UserID = user.ID
	result, err := c.Service.Create(r.Context(), data)
	helper.Response(w, result, err)
}

func (c *TweetController) FindByUserID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID := params.ByName("id")

	user, err := c.findUser(r.Context(), userID)
	if err != nil {
		helper.WriteErrorResponse(w, "User not found: "+err.Error(), http.StatusNotFound)
		return
	}

	result, err := c.Service.FindByUserID(r.Context(), user.ID)
	helper.Response(w, result, err)
}

func (c *TweetController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	err := c.Service.Delete(r.Context(), id)
	if err != nil {
		helper.WriteErrorResponse(w, "Error deleting tweet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	helper.Response(w, map[string]string{"message": "Tweet deleted successfully"}, nil)
}
