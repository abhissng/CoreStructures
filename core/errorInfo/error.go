package errorInfo

type ErrorInfo struct {
	Causes       []string       `json:"causes"`
	Component    string         `json:"component"`
	Description  string         `json:"description"`
	ErrorCode    string         `json:"errorCode"`
	Fields       map[string]any `json:"fields"`
	Message      string         `json:"message"`
	ResponseType string         `json:"responseType"`
	StatusCode   string         `json:"statusCode"`
}

func NewErrorInfo() *ErrorInfo {
	return &ErrorInfo{}
}
