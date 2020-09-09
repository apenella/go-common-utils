package errors

import (
	"fmt"
)

// Error is an error interface implementation that has an error context, an error message and could wrap other errors
type Error struct {
	context       string
	message       string
	wrappedErrors []error
}

// New returns an Error instance
func New(context, message string, errors ...error) *Error {
	return &Error{
		context, message, errors,
	}
}

// Error returns an error message with the wrapped errors
func (err *Error) Error() string {
	return fmt.Sprint(err.message, err.getWrappedErrors())
}

// ErrorWithContext returns an error message with its context and also the wrapped errors
func (err *Error) ErrorWithContext() string {
	if err.context != "" {
		return fmt.Sprintf("[%s] %s%s", err.context, err.message, err.getWrappedErrorsWithContext())
	} else {
		return fmt.Sprint(err.message, err.getWrappedErrorsWithContext())
	}
}

// getWrappedErrors
func (err *Error) getWrappedErrors() string {
	errors := ""

	for _, e := range err.wrappedErrors {
		errors = fmt.Sprintf("%s\n\t%s", errors, e.Error())
	}

	return errors
}

func (err *Error) getWrappedErrorsWithContext() string {

	errors := ""
	for _, e := range err.wrappedErrors {
		errorKind, ok := e.(*Error)
		if ok {
			errors = fmt.Sprintf("%s\n\t%s", errors, errorKind.ErrorWithContext())
		} else {
			errors = fmt.Sprintf("%s\n\t%s", errors, e.Error())
		}
	}
	return errors
}
