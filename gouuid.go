package gouuid

import (
	"github.com/google/uuid"
)

var isFreeze bool
var uuidFreeze string

func Freeze(id string) {
	uuidFreeze = id
	isFreeze = true
}

func Gen() string {
	if isFreeze {
		return uuidFreeze
	}

	return uuid.New().String()
}
