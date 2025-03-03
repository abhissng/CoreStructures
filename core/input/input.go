package input

// InputInfo represents the information about the input
type InputInfo struct {
	Action  string `json:"action,omitempty"`
	Subject string `json:"subject,omitempty"`
}

// NewInputInfo creates a new InputInfo
func NewInputInfo() *InputInfo {
	return &InputInfo{}
}
