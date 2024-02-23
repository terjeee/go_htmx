package tools

import (
	log "github.com/sirupsen/log"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

// TODO: hva er forskjellen på struct og interface
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var error error = database.SetupDatabase()
	if error != nil {
		log.Error(error)
		return nil, error
	}

	return &database, nil
}
