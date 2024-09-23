package v1

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/desulaidovich/auth/auth/domain"
	"github.com/desulaidovich/auth/auth/middleware"
	"github.com/desulaidovich/auth/auth/repository"
	"github.com/desulaidovich/auth/auth/usecase"
	"github.com/desulaidovich/auth/internal/ip"
	"github.com/desulaidovich/auth/internal/render"
	"github.com/desulaidovich/auth/internal/token"
	"github.com/jmoiron/sqlx"
)

type refreshHandler struct {
	usecase usecase.RefresTokenhUseCase
}

func NewRefreshHandler(db *sqlx.DB, k []byte) *refreshHandler {
	repo := repository.NewRepositoryPostgres(db)
	return &refreshHandler{
		usecase: *usecase.NewRefreshTokenUseCase(repo, k),
	}
}

func (h *refreshHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	id := middleware.RequestID(r.Context())
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})).With("request_id", id)

	req, err := render.Bind[domain.Request](r)
	if err != nil {
		if err := render.Render(&domain.Response{
			Error: "GUID was not found",
		}, http.StatusBadRequest, w); err != nil {
			logger.Debug("Internal error",
				slog.String("Message", err.Error()),
			)
		}
		return
	}

	refreshToken := r.Header.Get("Authorization")
	if refreshToken == "" {
		if err := render.Render(&domain.Response{
			Error: "Token was not found",
		}, http.StatusNotAcceptable, w); err != nil {
			logger.Debug("Internal error",
				slog.String("Message", err.Error()),
			)
		}
		return
	}

	newToken, err := h.usecase.RefreshToken(refreshToken, req.GUID, ip.IP(r))
	if err != nil {
		if errors.Is(err, usecase.ErrIPNotAllowed) {
			if err := render.Render(&domain.Response{
				Error: "Типа тут отправить email",
			}, http.StatusNotAcceptable, w); err != nil {
				logger.Debug("Internal error",
					slog.String("Message", err.Error()),
				)
			}
			return
		}

		if err := render.Render(&domain.Response{
			Error: "Token can't be refreshed",
		}, http.StatusBadRequest, w); err != nil {
			logger.Debug("Internal error",
				slog.String("Message", err.Error()),
			)
		}
		return
	}

	if err := render.Render(&domain.Response{
		Message: "New token",
		Token: &map[string]string{
			token.AccessToken: newToken,
		},
	}, http.StatusOK, w); err != nil {
		logger.Debug("Internal error",
			slog.String("Message", err.Error()),
		)
	}
}
