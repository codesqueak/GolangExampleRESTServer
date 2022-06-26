package model

import (
	"github.com/google/uuid"
)

// swagger:model User
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:model Users
type AllUsers struct {
	UserMap map[uuid.UUID]User `json:"users"`
}

// swagger:model Key
type Key struct {
	Key uuid.UUID `json:"key"`
}

// swagger:model ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}
