package render

import (
	"encoding/json"
	"io"
	"net/http"
)

func Bind[T any](r *http.Request) (*T, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	data := new(T)

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func Render[T any](t *T, statusCode int, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data, err := renderer(t)
	if err != nil {
		return err
	}

	w.Write(data)
	return nil
}

func renderer[T any](t *T) ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
