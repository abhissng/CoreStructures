package core

import (
	"github.com/abhissng/core-structures/core/primary"
	"github.com/abhissng/core-structures/core/user"
)

// Core represents the core structure of the application
type Core struct {
	User    *user.UserInformation `json:"user,omitempty"`
	Primary *primary.PrimaryInfo  `json:"primary,omitempty"`
}
