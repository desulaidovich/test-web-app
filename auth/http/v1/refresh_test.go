package v1_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/desulaidovich/auth/auth/domain"
	"github.com/desulaidovich/auth/internal/token"
)

// go test -run TestRefreshToken ./auth/http/v1/refresh_test.go -v
func TestRefreshToken(t *testing.T) {
	client := new(http.Client)

	requestBody, err := json.Marshal(&domain.Request{
		GUID: "123qwe",
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/refresh", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	defer request.Body.Close()

	request.Header.Add("Authorization", "REFRESH TOKEN HERE")
	request.Header.Add("X-Real-Ip", "IP ADDR HERE")

	response, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	responseBody := new(domain.Response)
	if err := json.Unmarshal(data, responseBody); err != nil {
		t.Fatal(err)
	}

	if responseBody.Error != "" {
		t.Fatalf("response: %s", responseBody.Error)
	}

	tokens := *responseBody.Token

	t.Logf("%s\n%s:%s\n%s:%s\n", responseBody.Message,
		token.AccessToken, tokens[token.AccessToken],
		token.RefreshToken, tokens[token.RefreshToken],
	)
}
