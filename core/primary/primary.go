package primary

// PrimaryInfo represents the information about the primary Information
type PrimaryInfo struct {
	Token string `json:"token,omitempty"`
}

// NewPrimaryInfo creates a new PrimaryInfo instance
func NewPrimaryInfo() *PrimaryInfo {
	return &PrimaryInfo{}
}
