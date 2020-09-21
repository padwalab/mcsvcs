package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/padwalab/mcsvcs/pkg/custom/endpoints"
)

// var ...
var (
	ErrUnknown = errors.New("unknown argument passed")

	ErrInvalidArgument = errors.New("invalid argument passed")
)

func NewHTTPHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()
	m.Handle("/test", httptransport.NewServer(
		ep.TestStatusEndpoint,
		decodeHTTPTestStatusRequest,
		encodeResponse,
	))
	return m
}

func decodeHTTPTestStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.TestStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
