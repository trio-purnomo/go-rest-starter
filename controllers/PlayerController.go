package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/trio-purnomo/go-rest-starter/helpers"
	"github.com/trio-purnomo/go-rest-starter/infrastructures"
	"github.com/trio-purnomo/go-rest-starter/models"
	"github.com/trio-purnomo/go-rest-starter/repositories"
	"github.com/trio-purnomo/go-rest-starter/services"
)

// InitPlayerController is
func InitPlayerController() *PlayerController {
	playerRepository := new(repositories.PlayerRepository)
	playerRepository.DB = &infrastructures.SQLConnection{}

	playerService := new(services.PlayerService)
	playerService.PlayerRepository = playerRepository

	playerController := new(PlayerController)
	playerController.PlayerService = playerService

	return playerController
}

// PlayerController is
type PlayerController struct {
	PlayerService services.IPlayerService
}

// StorePlayer is
func (p *PlayerController) StorePlayer(res http.ResponseWriter, req *http.Request) {
	var player models.Player
	//Read request data
	body, _ := ioutil.ReadAll(req.Body)
	err := json.Unmarshal(body, &player)

	if err != nil {
		helpers.Response(res, http.StatusBadRequest, "fail", "Failed read input data")
		return
	}

	result, err := p.PlayerService.StorePlayer(player)

	if err == nil {
		helpers.Response(res, http.StatusOK, "ok", result)
	} else {
		helpers.Response(res, http.StatusBadRequest, "fail", err.Error())
	}

	return
}

//GetPlayer is
func (p *PlayerController) GetPlayer(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		helpers.Response(res, http.StatusBadRequest, "fail", "Fail parse player id")
		return
	}

	player, err := p.PlayerService.GetPlayer(id)
	if err != nil {
		helpers.Response(res, http.StatusBadRequest, "fail", err.Error())
		return
	}

	if player.ID == 0 {
		helpers.Response(res, http.StatusOK, "ok", "Player not found")
		return
	}

	helpers.Response(res, http.StatusOK, "ok", player)
	return
}
