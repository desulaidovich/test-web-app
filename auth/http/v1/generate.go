package v1

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/desulaidovich/auth/auth/domain"
	"github.com/desulaidovich/auth/auth/middleware"
	"github.com/desulaidovich/auth/auth/repository"
	"github.com/desulaidovich/auth/auth/usecase"
	"github.com/desulaidovich/auth/internal/ip"
	"github.com/desulaidovich/auth/internal/render"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name=RenerageHandler
type RenerageHandler struct {
	usecase usecase.GenerateTokenUseCase
}

func NewGenerateHandler(db *sqlx.DB, k []byte) *RenerageHandler {
	repo := repository.NewRepositoryPostgres(db)
	return &RenerageHandler{
		usecase: *usecase.NewGenerateTokenUseCase(repo, k),
	}
}

func (h *RenerageHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	id := middleware.RequestID(r.Context())
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})).With("request_id", id)

	logger.Debug("Request start")
	defer logger.Debug("Request finish")

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

	tokens, err := h.usecase.GenerateToken(req.GUID, ip.IP(r))
	if err != nil {
		if err := render.Render(&domain.Response{
			Error: err.Error(),
		}, http.StatusInternalServerError, w); err != nil {
			logger.Debug("Internal error",
				slog.String("Message", err.Error()),
			)
		}
		return
	}

	if err := render.Render(&domain.Response{
		Message: "Welcome to the club, buddy",
		Token:   &tokens,
	}, http.StatusOK, w); err != nil {
		logger.Debug("Internal error",
			slog.String("Message", err.Error()),
		)
	}
}
