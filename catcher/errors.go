package catcher

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

// SdkError is a struct that implements the Go error interface.
type SdkError struct {
	Code  string         `json:"code"`
	Trace []string       `json:"trace"`
	Msg   string         `json:"msg"`
	Cause *string        `json:"cause,omitempty"`
	Args  map[string]any `json:"args,omitempty"`
}

// Error returns the error message.
func (e SdkError) Error() string {
	a, _ := json.Marshal(e)
	return string(a)
}

// Error tries to cast the cause as an SdkError, if it is not an SdkError, it creates a new SdkError with the given parameters.
// It logs the error message and returns the error.
// If cause is nil, it will store a blank string in the Cause field.
// The field Code is a hash of the message and trace. It is used to identify the recurrence of an error.
// Params:
// msg: the error message.
// cause: the error that caused this error.
// args: a map of additional information.
// Returns:
// *SdkError: the error. This type implements the Go error interface.
func Error(msg string, cause error, args map[string]any) *SdkError {
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

	var err *SdkError
	if err = ToSdkError(cause); err == nil {
		sum := md5.Sum([]byte(fmt.Sprint(msg, trace)))
		err = &SdkError{
			Code:  hex.EncodeToString(sum[:]),
			Trace: trace,
			Args:  args,
			Msg:   msg,
			Cause: func() *string {
				if cause != nil {
					return pointerOf(cause.Error())
				} else {
					return nil
				}
			}(),
		}
	}
	log.Println(err.Error())

	return err
}

// ToSdkError tries to cast an error to a SdkError.
// If the error is not an SdkError, it returns nil.
func ToSdkError(err error) *SdkError {
	if err == nil {
		return nil
	}

	var sdkError *SdkError
	switch {
	case errors.As(err, &sdkError):
		return err.(*SdkError)
	default:
		return nil
	}
}

// GinError is a helper function to return an error to the client using Gin framework context.
// It sets the headers x-error and x-error-id with the error message and UUID respectively and sets the status code.
func (e SdkError) GinError(c *gin.Context) {
	c.Header("x-error-id", e.Code)
	c.Header("x-error", e.Msg)

	status, ok := e.Args["status"].(int)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.AbortWithStatus(status)
	}
}

func pointerOf[t any](s t) *t {
	return &s
}