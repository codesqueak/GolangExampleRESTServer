package handlers

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"restServer/model"
	"strings"
)

// swagger:route GET /getuser/{uuid} User-API getUser
// Get a user
//
// produces
// - application/json
//
// Schemes: http
//
// Parameters:
//   + name: uuid
//     in: path
//     description: uuid allocated to the user record
//     required: true
//     type: string
//     format: uuid
//
// responses:
//  400: ErrorResponse Response error message
//  404: ErrorResponse No corresponding user
//  200: User Users record
//
func GetUserArgInURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["uuid"]
	if userId == "" {
		faultResponse(w, "Missing UUID", http.StatusBadRequest)
		return
	}
	userUuid, err := uuid.Parse(strings.TrimSpace(userId))
	if err != nil {
		faultResponseWithError(w, "Invalid UUID", http.StatusBadRequest, err)
		return
	}
	user, found := model.Users[userUuid]
	if found {
		writeJSONValue(w, user)
	} else {
		faultResponse(w, "User not found", http.StatusNotFound)
	}
}
