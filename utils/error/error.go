package error

// APIError Struct
type APIError struct {
	message string
}

func (error *APIError) Error() string {
	return error.message
}

// ThrowAPIError : Handler to throw API Errors
func ThrowAPIError(message string) {
	panic(&APIError{
		message: message,
	})
}
