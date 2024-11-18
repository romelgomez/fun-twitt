package account

import (
	"fmt"
	dto "funtwitt/domain/user/dto"
	service "funtwitt/domain/user/service"
	"funtwitt/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	Service service.UserService
}

func NewController(service service.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var data dto.UserCreate
	helper.ReadRequestBody(r, &data)
	if err := data.Validate(); err != nil {
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.Service.Create(r.Context(), data)
	helper.Response(w, result, err)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var data dto.UserUpdate
	helper.ReadRequestBody(r, &data)

	fmt.Println("")
	fmt.Println(data)
	fmt.Println("")

	result, err := c.Service.Update(r.Context(), data)
	helper.Response(w, result, err)
}

func (c *UserController) FindByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	result, err := c.Service.FindByID(r.Context(), id)

	helper.Response(w, result, err)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	result, err := c.Service.Delete(r.Context(), id)
	helper.Response(w, result, err)
}

func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result, err := c.Service.FindAll(r.Context())
	helper.Response(w, result, err)
}
