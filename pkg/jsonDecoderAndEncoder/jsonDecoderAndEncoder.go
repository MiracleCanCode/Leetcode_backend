package jsondecoderandencoder

import (
	"encoding/json"
	"errors"
	"net/http"
)

type JsonDecoderAndEncoder struct {
	r *http.Request
	w http.ResponseWriter
}

func NewJsonDecoderAndEncoder(r *http.Request, w http.ResponseWriter) *JsonDecoderAndEncoder {
	return &JsonDecoderAndEncoder{
		r: r,
		w: w,
	}
}

func (s *JsonDecoderAndEncoder) Decode(payload any) error {
	json := json.NewDecoder(s.r.Body)

	if err := json.Decode(&payload); err != nil {
		return errors.New("failed decode response")
	}

	return nil
}

func (s *JsonDecoderAndEncoder) Encode(payload any) error {
	json := json.NewEncoder(s.w)

	if err := json.Encode(&payload); err != nil {
		return errors.New("failed encode request")
	}

	return nil
}
