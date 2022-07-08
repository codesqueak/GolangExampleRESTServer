package handlers

import (
	"net/http"
	"restServer/model"
)

// swagger:route HEAD /getusers User-API getUser
// Get all user
//
// produces
// - application/json
//
// Schemes: http
//
// responses:
//  400: ErrorResponse Response error message
//  200: String Body will be empty
//
func GetAllUsersHeadHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSONValueHeadersOnly(w, model.AllUsers{model.Users})
}
