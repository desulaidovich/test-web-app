package usecase_test

import (
	"testing"

	"github.com/desulaidovich/auth/auth/mocks"
	"github.com/desulaidovich/auth/auth/usecase"
	"github.com/desulaidovich/auth/internal/token"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/mock"
)

// go test ./auth/usecase/generate_test.go -v
func TestGenerateTokenUseCase_GenerateToken(t *testing.T) {
	type args struct {
		guid string
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "all ok",
			args: args{
				guid: "asd123",
				addr: "localhost:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewRepository(t)

			repo.On("Add", tt.args.guid, mock.Anything).Return(nil, nil)

			u := usecase.NewGenerateTokenUseCase(repo, []byte("EXAMPLE"))

			d, err := u.GenerateToken(tt.args.guid, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTokenUseCase.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("Done:\n%s: %s\n%s: %s",
				token.AccessToken, d[token.AccessToken],
				token.RefreshToken, d[token.RefreshToken],
			)
		})
	}
}
