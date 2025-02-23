package input

// InputInfo represents the information about the input
type InputInfo struct {
	Action  string `json:"action"`
	Subject string `json:"subject"`
}

// NewInputInfo creates a new InputInfo
func NewInputInfo() *InputInfo {
	return &InputInfo{}
}
