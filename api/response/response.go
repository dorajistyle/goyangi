package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicResponse is a basic response.
type BasicResponse struct {
	Status      int    `json:"status",description:"http status"`
	MessageType string `json:"messageType", description:"response message type for frontend I18N."`
	Message     string `json:"message",description:"response message"`
}

// MessageTypes is an array of message types that using in frontend.
// OK                  : 200
// Created             : 201
// NotModified         : 304
// BadRequest          : 400
// Unauthorized        : 401
// PaymentRequired     : 402
// Forbidden           : 403
// NotFound            : 404
// MethodNotAllowed    : 405
// InternalServerError : 500
type MessageTypes struct {
	OK                  string `description:"200"`
	Created             string `description:"201"`
	NotModified         string `description:"304"`
	BadRequest          string `description:"400"`
	Unauthorized        string `description:"401"`
	PaymentRequired     string `description:"402"`
	Forbidden           string `description:"403"`
	NotFound            string `description:"404"`
	MethodNotAllowed    string `description:"405"`
	InternalServerError string `description:"500"`
}

// Messages is an array of messages
type Messages struct {
	OK      string `description:"200"`
	Created string `description:"201"`
}

// SuccessJSON is a json response of 2xx.
func SuccessJSON(c *gin.Context, status int, messageType string, msg string) {
	c.JSON(status, Success(status, messageType, msg))
}

// RedirectionJSON is a json response of 3xx.
func RedirectionJSON(c *gin.Context, status int, messageType string, msg string) {
	c.JSON(status, Redirection(status, messageType, msg))
}

// JSON is a json response.
func JSON(c *gin.Context, status int, messageTypes *MessageTypes, messages *Messages, err error) {
	if err == nil {
		SuccessJSON(c, status, messageTypes.OK, messages.OK)
	} else {
		ErrorJSON(c, status, messageTypes, err)
	}
}

// ErrorJSON is a json response of errors.
func ErrorJSON(c *gin.Context, status int, messageTypes *MessageTypes, err error) {
	switch status {
	case http.StatusNotModified:
		KnownErrorJSON(c, status, messageTypes.NotModified, err)
	case http.StatusBadRequest:
		KnownErrorJSON(c, status, messageTypes.BadRequest, err)
	case http.StatusUnauthorized:
		KnownErrorJSON(c, status, messageTypes.Unauthorized, err)
	case http.StatusPaymentRequired:
		KnownErrorJSON(c, status, messageTypes.PaymentRequired, err)
	case http.StatusForbidden:
		KnownErrorJSON(c, status, messageTypes.Forbidden, err)
	case http.StatusNotFound:
		KnownErrorJSON(c, status, messageTypes.NotFound, err)
	case http.StatusMethodNotAllowed:
		KnownErrorJSON(c, status, messageTypes.MethodNotAllowed, err)
	case http.StatusInternalServerError:
		KnownErrorJSON(c, status, messageTypes.InternalServerError, err)
	default:
		UnknownErrorJSON(c, status, err)
	}
}

// KnownErrorJSON is json response of known error.
func KnownErrorJSON(c *gin.Context, status int, messageType string, err error) {
	c.JSON(status, KnownError(status, messageType, err))
}

// UnknownErrorJSON is json response of unknown error.
func UnknownErrorJSON(c *gin.Context, status int, err error) {
	c.JSON(status, UnknownError(status, err))
}

// Success is a basic response of 2xx.
func Success(status int, messageType string, msg string) *BasicResponse {
	return &BasicResponse{Status: status, MessageType: messageType, Message: msg}
}

// Redirection is a basic response of 3xx.
func Redirection(status int, messageType string, msg string) *BasicResponse {
	return &BasicResponse{Status: status, MessageType: messageType, Message: msg}
}

// KnownError is a basic response of know errors.
func KnownError(status int, messageType string, err error) *BasicResponse {
	return &BasicResponse{Status: status, MessageType: messageType, Message: err.Error()}
}

// UnknownError is a basic response of unknown errors.
func UnknownError(status int, err error) *BasicResponse {
	return KnownError(status, "error.unknown", err)
}
