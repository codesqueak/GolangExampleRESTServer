package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"restServer/model"
)

// swagger:route POST /adduser User-API addUser
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
func AddUserPostHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
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
	if user.Username != "" {
		newUUID, err := uuid.NewRandom()
		if err != nil {
			faultResponseWithError(w, "Unable to generate UUID", http.StatusInternalServerError, err)
		} else {
			model.Users[newUUID] = user
			key := model.Key{Key: newUUID}
			writeJSONValue(w, key)
		}
	} else {
		faultResponse(w, "Invalid user name", http.StatusBadRequest)
	}
}
