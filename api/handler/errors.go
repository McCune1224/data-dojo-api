package handler

type ErrorResponse struct {
	Error             string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

const (
	Error400String = "bad_request"
	Error404String = "not_found"
	error500String = "internal_server_error"
)

var (
	// Used for returning a 400 status JSON response with error details
	error400Response = ErrorResponse{
		Error:             "bad_request",
		ErrorDescription: "The request was malformed",
	}

	// Used for returning a 404 status JSON response with error details
	Error404Response = ErrorResponse{
		Error:             "not_found",
		ErrorDescription: "The requested resource was not found",
	}

	// Used for returning a 500 status JSON response with error details
	Error500Response = ErrorResponse{
		Error:             "internal_server_error",
		ErrorDescription: "An internal server error occurred",
	}
)
