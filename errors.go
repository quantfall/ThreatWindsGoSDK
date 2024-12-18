package go_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

// SdkError is a struct that implements the Go error interface.
type SdkError struct {
	Code  string                 `json:"code"`
	Trace []string               `json:"trace,omitempty"`
	Args  map[string]interface{} `json:"args,omitempty"`
}

// Error returns the error message.
func (e SdkError) Error() string {
	a, _ := json.Marshal(e)
	return string(a)
}

// Error creates a new SdkError with a unique code generated from the trace and args.
func Error(trace []string, args map[string]interface{}) *SdkError {
	sum := md5.Sum([]byte(fmt.Sprint(trace, args)))

	return &SdkError{
		Code:  hex.EncodeToString(sum[:]),
		Trace: trace,
		Args:  args,
	}
}

// FromError converts a Go error with a JSON message to a SdkError.
// If the error message is not a valid JSON, it returns an error with the original message.
// If the error is nil, it returns nil.
// The error message should be a JSON object with the following fields:
// - code: string
// - trace: []string
// - args: map[string]interface{}
func FromError(err error) *SdkError {
	if err == nil {
		return nil
	}

	var e = new(SdkError)
	unmarshalErr := json.Unmarshal([]byte(err.Error()), e)
	if unmarshalErr != nil {
		return Error(Trace(), map[string]interface{}{
			"cause":         unmarshalErr.Error(),
			"error":         "could not parse error",
			"advise":        "please report this error to the developers",
			"originalError": err.Error(),
		})
	}

	return e
}

// Trace returns the stack trace of the caller.
func Trace() []string {
	pc := make([]uintptr, 25)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	var trace = make([]string, 0, 10)
	for {
		frame, more := frames.Next()

		trace = append(trace, fmt.Sprint(frame.Function, " ", frame.Line))
		if !more {
			break
		}
	}

	return trace
}

// GinError is a helper function to return an error to the client using Gin framework context.
// It sets the headers x-error and x-error-id with the error message and UUID respectively and sets the status code.
func (e SdkError) GinError(c *gin.Context) {
	c.Header("x-error-id", e.Code)

	err, ok := e.Args["error"].(string)
	if !ok {
		c.Header("x-error", "internal server error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.Header("x-error", err)
	}

	status, ok := e.Args["status"].(int)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.AbortWithStatus(status)
	}
}
