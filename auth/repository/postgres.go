package repository

import (
	"github.com/desulaidovich/auth/auth/models"
	"github.com/jmoiron/sqlx"
)

type TokenRepository struct {
	*sqlx.DB
}

func NewRepositoryPostgres(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{
		DB: db,
	}
}

func (r *TokenRepository) Add(guid, token string) (*models.Token, error) {
	t := &models.Token{
		GUID:         guid,
		RegreshToken: []byte(token),
	}

	rows, err := r.DB.NamedQuery(`INSERT INTO public.token
		(guid, refresh_token) VALUES (:guid, :refresh_token) RETURNING *;`, t)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.StructScan(&t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (r *TokenRepository) Get(guid string) (*models.Token, error) {
	t := &models.Token{
		GUID: guid,
	}

	rows, err := r.DB.NamedQuery(`SELECT * FROM public.token
		WHERE guid=:guid;`, &t)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.StructScan(&t); err != nil {
			return nil, err
		}
	}

	return t, nil
}
