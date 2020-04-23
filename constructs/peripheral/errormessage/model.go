package errormessage

import (
	"time"
)

// ErrorMessage ...
type ErrorMessage struct {
	ErrorType  string    `json:"errorType"`
	Message    string    `json:"message"`
	OccurredAt time.Time `json:"occurredAt" format:"date-time"`
}

// NewError ...
func NewError(errortype string, message string) ErrorMessage {
	return ErrorMessage{
		ErrorType:  errortype,
		Message:    message,
		OccurredAt: time.Now(),
	}
}
