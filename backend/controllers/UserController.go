package controllers

import (
	"guess_the_score/backend/models"
	"guess_the_score/backend/views"
	"net/http"
)

func NewUserController(us models.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

type UserController struct {
	us models.UserService
}

func (self *UserController) GetTop(w http.ResponseWriter, r *http.Request) {

	views.SendResponse(w, "privet", http.StatusOK)
}