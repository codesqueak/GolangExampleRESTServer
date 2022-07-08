package handlers

import (
	"github.com/go-http-utils/headers"
	"net/http"
)

// swagger:route OPTIONS /users User-API getUser
// Get available server OPTION information
//
// produces
// - application/json
//
// Schemes: http
//
// responses:
//  200: String Body will be empty
//
func OptionsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(headers.ContentType, "application/json; charset=utf-8")
	w.Header().Set(headers.Allow, "GET,HEAD,POST,PUT,DELETE,OPTIONS")
	w.WriteHeader(http.StatusOK)
}
