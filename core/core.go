package core

import (
	"github.com/abhissng/core-structures/core/input"
	"github.com/abhissng/core-structures/core/primary"
	"github.com/abhissng/core-structures/core/user"
)

// Core represents the core structure of the application
type Core struct {
	User    *user.UserInformation `json:"user,omitempty"`
	Primary *primary.PrimaryInfo  `json:"primary,omitempty"`
	Input   *input.InputInfo      `json:"input,omitempty"`
	Error   any                   `json:"error,omitempty"`
}

func NewCore() *Core {
	return &Core{
		User:    user.NewUserInformation(),
		Primary: primary.NewPrimaryInfo(),
		Input:   input.NewInputInfo(),
	}
}
func (c *Core) AttachError(err any) *Core {
	c.Error = err
	return c
}
