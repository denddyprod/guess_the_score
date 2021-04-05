package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guess_the_score/backend/models"
	"guess_the_score/backend/views"
	"log"
	"net/http"
)

func NewAuthController(us models.UserService) *AuthController {
	return &AuthController{
		us: us,
	}
}

type AuthController struct {
	us models.UserService
}

func (self *AuthController) Index(w http.ResponseWriter, r *http.Request) {
	userId, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}
	//decoder := json.NewDecoder(r.Body)
	//var bodyUser *models.User
	//err := decoder.Decode(&bodyUser)
	//if err != nil {
	//	log.Println(err)
	//	views.SendResponse(w, nil, http.StatusForbidden)
	//	return
	//}

	resUser, err := self.us.FindById(userId)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	views.SendResponse(w, resUser, http.StatusCreated)
	return
}

func (self *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var loginUser *models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginUser)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	err = self.us.Create(loginUser)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	signupRes := map[string]string{"success": "true"}
	views.SendResponse(w, signupRes, http.StatusCreated)
	return
}

func (self *AuthController) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var bodyUser *models.User
	err := decoder.Decode(&bodyUser)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	log.Println(bodyUser)

	err = self.us.Update(bodyUser)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	signupRes := map[string]string{"success": "true"}
	views.SendResponse(w, signupRes, http.StatusCreated)
	return
}

func (self *AuthController) Delete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var bodyUser *models.User
	err := decoder.Decode(&bodyUser)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	err = self.us.Delete(bodyUser.Id)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	signupRes := map[string]string{"success": "true"}
	views.SendResponse(w, signupRes, http.StatusCreated)
	return
}
