package usecase

import (
	"errors"
	"log"

	"github.com/desulaidovich/auth/auth"
	"github.com/desulaidovich/auth/internal/token"
)

type GenerateTokenUseCase struct {
	repo      auth.Repository
	secretKey []byte
}

func NewGenerateTokenUseCase(r auth.Repository, k []byte) *GenerateTokenUseCase {
	return &GenerateTokenUseCase{
		repo:      r,
		secretKey: k,
	}
}

func (u *GenerateTokenUseCase) GenerateToken(guid, addr string) (map[string]string, error) {
	if guid == "" {
		return nil, errors.New("guid is empty")
	}

	if addr == "" {
		return nil, errors.New("guid is empty")
	}

	accessToken, err := token.CreateAccess(guid, addr)
	if err != nil {
		return nil, err
	}

	accessTokenStr, err := accessToken.SignedString(u.secretKey)
	if err != nil {
		return nil, err
	}

	refreshTokenStr, err := token.CreateRefresh(guid, addr, u.secretKey)
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		token.AccessToken:  accessTokenStr,
		token.RefreshToken: refreshTokenStr,
	}

	if _, err := u.repo.Add(guid, refreshTokenStr); err != nil {
		log.Printf("can't insert data: %v", err)
		return nil, err
	}

	return data, nil
}
