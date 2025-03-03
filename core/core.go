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
}

// NewCore creates a new Core instance
func NewCore() *Core {
	return &Core{
		User:    user.NewUserInformation(),
		Primary: primary.NewPrimaryInfo(),
		Input:   input.NewInputInfo(),
	}
}

// NewNilCore creates a new Core instance
func NewNilCore() *Core {
	return &Core{}
}

// AttachUser attaches the given user information to the core
func (c *Core) AttachUser(user *user.UserInformation) *Core {
	c.User = user
	return c
}

// AttachPrimary attaches the given primary information to the core
func (c *Core) AttachPrimary(primary *primary.PrimaryInfo) *Core {
	c.Primary = primary
	return c
}

// AttachInput attaches the given input information to the core
func (c *Core) AttachInput(input *input.InputInfo) *Core {
	c.Input = input
	return c
}
