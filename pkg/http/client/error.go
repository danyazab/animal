package client

type (
	// HttpClientError represents HTTP error returned by airSlate client.
	HttpClientError struct {
		msg  string
		code int
	}

	// GenericHttpError represents generic interface that suits many packages.
	GenericHttpError interface {
		String() string
		StatusCode() int
	}
)

// NewHTTPClientError creates a new HttpClientError instance using
// generic interface that suits many packages.
func NewHTTPClientError(resp GenericHttpError) *HttpClientError {
	return &HttpClientError{
		msg:  resp.String(),
		code: resp.StatusCode(),
	}
}

// Code returns HTTP status code.
func (e *HttpClientError) Code() int {
	return e.code
}

func (e *HttpClientError) Error() string {
	return e.msg
}
