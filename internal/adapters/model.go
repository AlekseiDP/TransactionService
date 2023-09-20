package adapters

type ApiResult struct {
	Data         any    `json:"data,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorDisplay string `json:"errorDisplay,omitempty"`
}
