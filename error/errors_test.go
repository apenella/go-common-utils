package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		desc    string
		context string
		message string
		errs    []error
		res     *Error
	}{
		{
			desc:    "Testing new error",
			context: "context",
			message: "message",
			errs: []error{
				fmt.Errorf("wrapped error"),
				&Error{
					context: "wrapped error context",
					message: "wrapped error message",
				},
			},
			res: &Error{
				context: "context",
				message: "message",
				wrappedErrors: []error{
					fmt.Errorf("wrapped error"),
					&Error{
						context: "wrapped error context",
						message: "wrapped error message",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := New(test.context, test.message, test.errs...)
			assert.Equal(t, test.res, res, "Unexpected value")
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		desc string
		err  *Error
		res  string
	}{
		{
			desc: "Testing simple error with no context",
			err: &Error{
				message: "that is an error message",
			},
			res: "that is an error message",
		},
		{
			desc: "Testing simple error with context",
			err: &Error{
				context: "context is not shown",
				message: "that is an error message",
			},
			res: "that is an error message",
		},
		{
			desc: "Testing simple error with wrapped errors",
			err: &Error{
				context: "context is not shown",
				message: "that is an error message",
				wrappedErrors: []error{
					fmt.Errorf("another error"),
					&Error{
						message: "even another error",
					},
				},
			},
			res: "that is an error message\n\tanother error\n\teven another error",
		},
		{
			desc: "Testing simple error with wrapped errors with nested errors",
			err: &Error{
				context: "context is not shown",
				message: "that is an error message",
				wrappedErrors: []error{
					fmt.Errorf("another error"),
					&Error{
						message: "even another error",
						wrappedErrors: []error{
							&Error{
								message: "even even another error",
							},
						},
					},
				},
			},
			res: "that is an error message\n\tanother error\n\teven another error\n\teven even another error",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := test.err.Error()
			assert.Equal(t, test.res, res, "Unexpected error message")
		})
	}
}

func TestErrorWithContext(t *testing.T) {
	tests := []struct {
		desc string
		err  *Error
		res  string
	}{
		{
			desc: "Testing simple error with no context",
			err: &Error{
				message: "that is an error message",
			},
			res: "that is an error message",
		},
		{
			desc: "Testing simple error with context",
			err: &Error{
				context: "Testing context",
				message: "that is an error message",
			},
			res: "[Testing context] that is an error message",
		},
		{
			desc: "Testing simple error with wrapped errors",
			err: &Error{
				context: "Testing context",
				message: "that is an error message",
				wrappedErrors: []error{
					fmt.Errorf("another error"),
					&Error{
						context: "Testing even context",
						message: "even another error",
					},
				},
			},
			res: "[Testing context] that is an error message\n\tanother error\n\t[Testing even context] even another error",
		},
		{
			desc: "Testing simple error with wrapped errors with nested errors",
			err: &Error{
				context: "Testing context",
				message: "that is an error message",
				wrappedErrors: []error{
					fmt.Errorf("another error"),
					&Error{
						context: "Testing even context",
						message: "even another error",
						wrappedErrors: []error{
							&Error{
								context: "Testing even even context",
								message: "even even another error",
							},
						},
					},
				},
			},
			res: "[Testing context] that is an error message\n\tanother error\n\t[Testing even context] even another error\n\t[Testing even even context] even even another error",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := test.err.ErrorWithContext()
			assert.Equal(t, test.res, res, "Unexpected error message")
		})
	}
}
