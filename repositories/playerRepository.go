package repositories

import (

	//"github.com/afex/hystrix-go/hystrix"

	"github.com/afex/hystrix-go/hystrix"
	log "github.com/sirupsen/logrus"
	"github.com/trio-purnomo/go-rest-starter/infrastructures"
	"github.com/trio-purnomo/go-rest-starter/models"
)

// IPlayerRepository is
type IPlayerRepository interface {
	StorePlayer(data models.Player) (models.Player, error)
}

// PlayerRepository is
type PlayerRepository struct {
	DB infrastructures.ISQLConnection
}

// StorePlayer store agent type data to database
func (r *PlayerRepository) StorePlayer(data models.Player) (models.Player, error) {
	err := hystrix.Do("StorePlayer", func() error {
		db := r.DB.GetPlayerWriteDb()
		defer db.Close()

		stmt, err := db.Prepare(`INSERT INTO players(players.name, players.score) VALUES(?, ?)`)
		if err != nil {
			return err
		}

		res, err := stmt.Exec(
			data.Name,
			data.Score,
		)

		if err != nil {
			return err
		}

		data.ID, err = res.LastInsertId()
		return err
	}, nil)

	if err != nil {
		log.WithFields(log.Fields{
			"event": "StorePlayer",
			"data":  data,
		}).Info(err)
	}

	return data, err
}
