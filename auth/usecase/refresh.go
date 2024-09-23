package usecase

import (
	"errors"

	"github.com/desulaidovich/auth/auth"
	"github.com/desulaidovich/auth/internal/token"
	"github.com/golang-jwt/jwt/v5"
)

var ErrIPNotAllowed = errors.New("SEND EMAIL AAAAAAAAAAA")

type RefresTokenhUseCase struct {
	repo      auth.Repository
	secretKey []byte
}

func NewRefreshTokenUseCase(r auth.Repository, k []byte) *RefresTokenhUseCase {
	return &RefresTokenhUseCase{
		repo:      r,
		secretKey: k,
	}
}

func (u *RefresTokenhUseCase) RefreshToken(refresh, guid, addr string) (string, error) {
	if _, err := u.repo.Get(refresh); err != nil {
		return "", errors.New("token not found")
	}

	refreshTokenStr, err := token.Decode(refresh)
	if err != nil {
		return "", err
	}

	refreshToken, err := token.Parse(refreshTokenStr, func(t *jwt.Token) (any, error) {
		return u.secretKey, nil
	})
	if err != nil {
		return "", err
	}

	iis, err := refreshToken.GetIssuer()
	if err != nil {
		return "", err
	}

	sub, err := refreshToken.GetSubject()
	if err != nil {
		return "", err
	}

	if iis != addr || sub != guid {
		return "", ErrIPNotAllowed
	}

	newAccessToken, err := token.CreateAccess(guid, addr)
	if err != nil {
		return "", err
	}

	newAccessTokenStr, err := newAccessToken.SignedString(u.secretKey)
	if err != nil {
		return "", err
	}

	return newAccessTokenStr, nil
}
