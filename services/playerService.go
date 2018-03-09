package services

import (
	"errors"

	"github.com/trio-purnomo/go-rest-starter/models"
	"github.com/trio-purnomo/go-rest-starter/repositories"
)

// IPlayerService is
type IPlayerService interface {
	StorePlayer(models.Player) (models.Player, error)
}

// PlayerService is
type PlayerService struct {
	PlayerRepository repositories.IPlayerRepository
}

// StorePlayer is
func (p *PlayerService) StorePlayer(data models.Player) (result models.Player, err error) {
	result, err = p.PlayerRepository.StorePlayer(data)
	if err != nil {
		err = errors.New("Failed save data to database")
	}
	return result, err
}
