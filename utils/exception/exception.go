package exception

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Code      int         `json:"-"`       // HTTP status code (for Fiber only)
	Message   string      `json:"message"` // User-friendly error message
	Status    bool        `json:"status"`  // false for error
	ErrorCode string      `json:"code"`    // App-level code (e.g., VALIDATION_ERROR)
	Data      interface{} `json:"data"`    // Additional data (e.g., map of validation errors)
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("code=%d message=%s", e.Code, e.Message)
}

// ===== GENERIC ERROR BUILDERS =====

func NewHttpError(code int, message string, errorCode string, data interface{}) *HttpError {
	return &HttpError{
		Code:      code,
		Message:   message,
		Status:    false,
		ErrorCode: errorCode,
		Data:      data,
	}
}

// ===== COMMON HTTP ERROR HELPERS =====

func BadRequest(msg string) *HttpError {
	return NewHttpError(http.StatusBadRequest, msg, "BAD_REQUEST", nil)
}

func BadRequestWithData(msg string, data interface{}) *HttpError {
	return NewHttpError(http.StatusBadRequest, msg, "BAD_REQUEST", data)
}

func NotFound(msg string) *HttpError {
	return NewHttpError(http.StatusNotFound, msg, "NOT_FOUND", nil)
}

func Unauthorized(msg string) *HttpError {
	return NewHttpError(http.StatusUnauthorized, msg, "UNAUTHORIZED", nil)
}

func Forbidden(msg string) *HttpError {
	return NewHttpError(http.StatusForbidden, msg, "FORBIDDEN", nil)
}

func Conflict(msg string) *HttpError {
	return NewHttpError(http.StatusConflict, msg, "CONFLICT", nil)
}

func TooManyRequests(msg string) *HttpError {
	return NewHttpError(http.StatusTooManyRequests, msg, "TOO_MANY_REQUESTS", nil)
}

func UnprocessableEntity(msg string) *HttpError {
	return NewHttpError(http.StatusUnprocessableEntity, msg, "UNPROCESSABLE_ENTITY", nil)
}

func UnprocessableEntityWithData(msg string, data interface{}) *HttpError {
	return NewHttpError(http.StatusUnprocessableEntity, msg, "UNPROCESSABLE_ENTITY", data)
}

func InternalServerError(msg string) *HttpError {
	return NewHttpError(http.StatusInternalServerError, msg, "INTERNAL_ERROR", nil)
}

func ServiceUnavailable(msg string) *HttpError {
	return NewHttpError(http.StatusServiceUnavailable, msg, "SERVICE_UNAVAILABLE", nil)
}

func GatewayTimeout(msg string) *HttpError {
	return NewHttpError(http.StatusGatewayTimeout, msg, "GATEWAY_TIMEOUT", nil)
}
