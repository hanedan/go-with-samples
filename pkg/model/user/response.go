package model

const (
	Success = "success"
	Error   = "error"
)

const (
	ErrEmptyRequestBody       = "request body is emtpy"
	ErrInvalidRequestMethod   = "invalid request method"
	ErrRequestBodyParseFailed = "couldn't parse request body: "
	ErrValidationFailed       = "validation failed for request: "
	ErrCouldntCreateUser      = "couldn't create user: "
)

type Response struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`

	// Valid only for Error responses
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}

// Options Pattern
// Option function to set the status and status code
type Option func(r *Response)

// Option function to set the error code and error message
func WithError(code int, message string) Option {
	return func(r *Response) {
		r.Status = Error
		r.ErrorCode = code
		r.Error = message
	}
}

// Option function to set the success status and status code
func WithSuccess(code int) Option {
	return func(r *Response) {
		r.Status = Success
		r.StatusCode = code
	}
}

// Function to create a new response with options
func NewResponse(options ...Option) *Response {
	response := &Response{}
	for _, option := range options {
		option(response)
	}
	return response
}
