package token

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
)

func CreateAccess(guid, addr string) (*jwt.Token, error) {
	payload := jwt.MapClaims{
		"sub": guid,
		"iss": addr,
		"exp": jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token, nil
}

func Parse(input string, keyFunc jwt.Keyfunc) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(input, jwt.MapClaims{}, keyFunc)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	data, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token data")
	}

	return &data, nil
}

func CreateRefresh(guid, addr string, key []byte) (string, error) {
	payload := jwt.MapClaims{
		"sub": guid,
		"iss": addr,
		"exp": jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(t)), nil
}

func Decode(token string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
