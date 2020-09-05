package error

// APIError Struct
type APIError struct {
	message string
}

func (error *APIError) Error() string {
	return error.message
}

// GetAPIError : return APIError
func GetAPIError(message string) *APIError {
	return &APIError{
		message: message,
	}
}

// ThrowAPIError : Handler to throw API Errors
func ThrowAPIError(message string) {
	panic(GetAPIError(message))
}
