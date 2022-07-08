package handlers

import (
	"net/http"
	"restServer/model"
)

// swagger:route GET /getusers User-API getUser
// Get all user
//
// produces
// - application/json
//
// Schemes: http
//
// responses:
//  400: ErrorResponse Response error message
//  200: Users All user records
//
func GetAllUsersHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSONValue(w, model.AllUsers{model.Users})
}
