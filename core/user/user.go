package user

// UserInformation represents information about a user
type UserInformation struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

// NewUserInformation creates a new UserInformation
func NewUserInformation() *UserInformation {
	return &UserInformation{}
}
