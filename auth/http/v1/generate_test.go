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

// go test -run TestGenerateToken ./auth/http/v1/generate_test.go -v
func TestGenerateToken(t *testing.T) {
	client := new(http.Client)

	requestBody, err := json.Marshal(&domain.Request{
		GUID: "123qwe",
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/generate", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	defer request.Body.Close()

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
