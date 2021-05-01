package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"guess_the_score/backend/models"
	"guess_the_score/backend/views"
	"log"
	"net/http"
)

func NewPredictionController(mm models.MatchModel, pm models.PredictionModel) *PredictionController {
	return &PredictionController{
		mm: mm,
		pm: pm,
	}
}

type PredictionController struct {
	mm models.MatchModel
	pm models.PredictionModel
}

func (self *PredictionController) GetAllByUserId(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	userId := params["id"]

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

	type response struct {
		TeamAName  string
		TeamBName  string
		TeamAScore string
		TeamBScore string
	}

	var resPredictions []response
	for _, prediction := range predictions {
		match, err := self.mm.FindById(prediction.MatchId)
		if err != nil {
			continue
		}

		var res response
		res.TeamAName = match.TeamA
		res.TeamBName = match.TeamB
		res.TeamAScore = prediction.TeamAScore
		res.TeamBScore = prediction.TeamBScore

		resPredictions = append(resPredictions, res)
	}

	views.SendResponse(w, resPredictions, http.StatusOK)
}

func (self *PredictionController) Create(w http.ResponseWriter, r *http.Request) {
	var prediction *models.Prediction

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&prediction)
	if err != nil {
		log.Println(err)
		views.SendResponse(w, nil, http.StatusForbidden)
		return
	}

	err = self.pm.Create(prediction)
	if err != nil {
		log.Println(err)
		data := map[string]string{"success": "false", "errorMsg": err.Error()}
		views.SendResponse(w, data, http.StatusForbidden)
		return
	}

	views.SendResponse(w, nil, http.StatusOK)
}
