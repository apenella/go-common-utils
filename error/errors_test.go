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
			t.Log(test.desc)

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
			res: "that is an error message\n another error\n even another error",
		},
		{
			desc: "Testing simple error with wrapped errors and nested errors",
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
			res: "that is an error message\n another error\n even another error\n even even another error",
		},
		{
			desc: "Testing propagation of errors",
			err: &Error{
				context: "Testing context",
				message: "",
				wrappedErrors: []error{
					&Error{
						message: "",
						wrappedErrors: []error{
							&Error{
								message: "",
								wrappedErrors: []error{
									fmt.Errorf("error propagated"),
								},
							},
						},
					},
				},
			},
			res: "error propagated",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

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
			res: "[Testing context] that is an error message\n another error\n [Testing even context] even another error",
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
			res: "[Testing context] that is an error message\n another error\n [Testing even context] even another error\n [Testing even even context] even even another error",
		},
		{
			desc: "Testing propagation of errors with context",
			err: &Error{
				context: "Testing context 1",
				message: "",
				wrappedErrors: []error{
					&Error{
						context: "Testing context 2",
						message: "",
						wrappedErrors: []error{
							&Error{
								message: "",
								context: "Testing context 3",
								wrappedErrors: []error{
									fmt.Errorf("error propagated 3"),
								},
							},
							&Error{
								message: "",
								context: "Testing context 4",
								wrappedErrors: []error{
									fmt.Errorf("error propagated 4"),
								},
							},
						},
					},
				},
			},
			res: "[Testing context 1] \n [Testing context 2] \n [Testing context 3] \n error propagated 3\n [Testing context 4] \n error propagated 4",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res := test.err.ErrorWithContext()
			assert.Equal(t, test.res, res, "Unexpected error message")
		})
	}
}
