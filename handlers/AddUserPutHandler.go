package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"restServer/model"
	"strings"
)

// swagger:route PUT /adduser/{uuid} User-API addUser
// Add a user
//
// consumes
// - application/json
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
//   + name: user
//     in: body
//     description: New user data
//     required: true
//     type: User
//
// responses:
//  400: ErrorResponse Response error message
//  200: Key Unique identifier for the user (UUID)
//
func AddUserPutHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	vars := mux.Vars(r)
	userId := vars["uuid"]
	if userId == "" {
		faultResponse(w, "Missing uuid parameter", http.StatusBadRequest)
		return
	}
	userUuid, err := uuid.Parse(strings.TrimSpace(userId))
	if err != nil {
		faultResponseWithError(w, "Invalid UUID", http.StatusBadRequest, err)
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		faultResponseWithError(w, "Unable to read message body", http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		faultResponseWithError(w, "Unable to process json", http.StatusBadRequest, err)
		return
	}
	if user.Username == "" {
		faultResponse(w, "Invalid user name", http.StatusBadRequest)
		return
	}
	originalUser, found := model.Users[userUuid]
	key := model.Key{Key: userUuid}
	if found {
		if user.Username == originalUser.Username && user.Password == originalUser.Password {
			writeJSONValue(w, key)
		} else {
			faultResponse(w, "Invalid attempt to change existing user record", http.StatusBadRequest)
		}
	} else {
		model.Users[userUuid] = user
		writeJSONValue(w, key)
	}

}
