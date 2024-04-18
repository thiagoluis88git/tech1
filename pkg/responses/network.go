package responses

import (
	"net/http"
	"net/url"
	"strings"
)

type NetworkError struct {
	Code    int
	Message string
}

func (er NetworkError) Error() string {
	return er.Message
}

func GetNetworkError(urlError *url.Error) *NetworkError {
	code := http.StatusInternalServerError
	message := urlError.Unwrap().Error()

	if urlError.Timeout() {
		code = http.StatusGatewayTimeout
	}

	if urlError.Temporary() {
		code = http.StatusServiceUnavailable
	}

	return &NetworkError{
		Code:    code,
		Message: message,
	}
}

func IsNetworkResponseOk(response *http.Response, bodyMessage string) error {
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}

	return &NetworkError{
		Code:    response.StatusCode,
		Message: bodyMessage,
	}
}

func AdjustStatusCodeFromMessageError(networError *NetworkError) error {
	statusCode := networError.Code

	if strings.Contains(networError.Message, "jwt expired") {
		statusCode = http.StatusUnauthorized
	}

	if strings.Contains(networError.Message, "jwt malformed") {
		statusCode = http.StatusBadRequest
	}

	if strings.Contains(networError.Message, "jwt must be provided") {
		statusCode = http.StatusBadRequest
	}

	if strings.Contains(networError.Message, "invalid token") {
		statusCode = http.StatusBadRequest
	}

	return &NetworkError{
		Code:    statusCode,
		Message: networError.Message,
	}
}
