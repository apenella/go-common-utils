package common

import (
	"fmt"
	"strings"
)

// StringList is type for a list of strings
type StringList []string

// String convers a stringlist to string
func (s *StringList) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set a value to stringlist
func (s *StringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}
