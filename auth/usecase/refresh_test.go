package usecase_test

import (
	"testing"

	"github.com/desulaidovich/auth/auth/mocks"
	"github.com/desulaidovich/auth/auth/usecase"
	"github.com/desulaidovich/auth/internal/token"
)

// go test ./auth/usecase/refresh_test.go -v
func TestRefresTokenhUseCase_RefreshToken(t *testing.T) {
	type args struct {
		refresh string
		guid    string
		addr    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				refresh: "ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmxlSEFpT2pFM01qY3lOall3Tnpnc0ltbHpjeUk2SW14dlkyRnNhRzl6ZERvNE1EZ3dJaXdpYzNWaUlqb2lZWE5rTVRJekluMC5tMlRVMENydFZlRFJtTENWa3p4ZUJ1NTRyYXRzVjhXTmZlTl9OalVrdnkw",
				guid:    "asd123",
				addr:    "localhost:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)

			repo.On("Get", tt.args.refresh).Return(nil, nil)

			u := usecase.NewRefreshTokenUseCase(repo, []byte("EXAMPLE"))

			d, err := u.RefreshToken(tt.args.refresh, tt.args.guid, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefresTokenUseCase.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("Done:\n%s: %s",
				token.AccessToken, d,
			)
		})
	}
}
