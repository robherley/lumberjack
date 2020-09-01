package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequestError holds the message string and http code
type RequestError struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	// ErrInvalidParam either means the given route parameter was wrong, like a non uint, or too long
	ErrInvalidParam = RequestError{Error: "Bad Request", Code: http.StatusBadRequest}
	// ErrInternalError is when a server error occurs, we do not want to send any info about the error back to the end user
	ErrInternalError = RequestError{Error: "Internal Error", Code: http.StatusInternalServerError, Message: "Internal Server Error. Please try again / contact an admin if this is an on going issue."}
	// ErrNotFound is when a resource or location is not found
	ErrNotFound = RequestError{Error: "Request Not Found", Code: http.StatusNotFound}
	// ErrUnauthorized means the user could not be validated and any JWT tokens on client side should be removed
	ErrUnauthorized = RequestError{Error: "Unauthorized", Code: http.StatusUnauthorized}
	// ErrForbidden is either anon accessing a route that requires auth, or an authed user without the correct permissions
	ErrForbidden = RequestError{Error: "Forbidden", Code: http.StatusForbidden}
	// ErrBadRequest the server cannot or will not process the request due to an apparent client error
	ErrBadRequest = RequestError{Error: "Bad Request", Code: http.StatusForbidden}
	// ErrConflict a request conflict with current state of the server
	ErrConflict = RequestError{Error: "Conflict", Code: http.StatusConflict}
)

// RespondWithError is a simple error handler that will log the error, form a proper
// response error object and send it
func RespondWithError(ctx *gin.Context, reqErr RequestError, messages ...string) {
	errWithMessage := reqErr
	if len(messages) != 0 {
		errWithMessage.Message = strings.Join(messages, " ")
	} else {
		errWithMessage.Message = "Unable To Complete Request"
	}

	ctx.JSON(errWithMessage.Code, errWithMessage)
}
