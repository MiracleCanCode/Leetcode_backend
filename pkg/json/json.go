package json

import (
	"encoding/json"
	"errors"
	"net/http"
)

type io struct {
	r *http.Request
	w http.ResponseWriter
}

func New(r *http.Request, w http.ResponseWriter) *io {
	return &io{
		r: r,
		w: w,
	}
}

func (s *io) Decode(payload any) error {
	json := json.NewDecoder(s.r.Body)

	if err := json.Decode(&payload); err != nil {
		return errors.New("failed decode response")
	}

	return nil
}

func (s *io) Encode(payload any) error {
	json := json.NewEncoder(s.w)

	if err := json.Encode(&payload); err != nil {
		return errors.New("failed encode request")
	}

	return nil
}
