package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Json utils

func ParseJsonBody(r *http.Request, v any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}

func JsonResponse(w http.ResponseWriter, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}
	return nil
}
