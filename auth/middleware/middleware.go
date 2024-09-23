package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type requestConextKey string

const requestID requestConextKey = "traceID"

func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()

		ctx := context.WithValue(r.Context(), requestID, id)

		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("request_id", id)
		logger.Info("Received init")

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func RequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestID).(string); ok {
		return id
	}

	return "empty request id"
}
