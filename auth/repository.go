package auth

import "github.com/desulaidovich/auth/auth/models"

//go:generate go run github.com/vektra/mockery/v2@v2.46.0 --name Repository
type Repository interface {
	Add(guid, token string) (*models.Token, error)
	Get(guid string) (*models.Token, error)
}
