package controllers

import (
	"fmt"
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

	topUsers, err := self.us.GetTop()
	if err != nil {
		fmt.Println(err)
		data := map[string]string{"success": "false", "errorMsg": err.Error()}
		views.SendResponse(w, data, http.StatusForbidden)
	}

	views.SendResponse(w, topUsers, http.StatusOK)
}
