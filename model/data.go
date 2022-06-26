package model

import (
	"github.com/google/uuid"
)

var Users = make(map[uuid.UUID]User)

var Port string
