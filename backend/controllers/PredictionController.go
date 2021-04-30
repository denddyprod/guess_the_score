package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guess_the_score/backend/models"
	"guess_the_score/backend/views"
	"net/http"
)

func NewPredictionController(pm models.PredictionModel) *PredictionController {
	return &PredictionController{
		pm: pm,
	}
}

type PredictionController struct {
	pm models.PredictionModel
}

func (self *PredictionController) GetAllByUserId(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	userId := params["id"]

	fmt.Println(userId)

	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println(err)
		data := map[string]string{"success": "false", "errorMsg": err.Error()}
		views.SendResponse(w, data, http.StatusForbidden)
	}

	predictions, err := self.pm.FindByUserId(objId)
	if err != nil {
		fmt.Println(err)
		data := map[string]string{"success": "false", "errorMsg": err.Error()}
		views.SendResponse(w, data, http.StatusForbidden)
		return
	}

	views.SendResponse(w, predictions, http.StatusOK)
}