// Copyright 2017 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package validation

import (
	"fmt"
	"strings"
)

// Validatable is used by entities that support validation.
type Validatable interface {
	Validate() Errors
}

// Error represents a single error in validation.
type Error struct {
	Field string `json:"field"` // Field representing the error, e.g. Email
	Issue string `json:"issue"` // Issue describing the error in detail
}

// String returns a textual representation of the validation error.
func (e Error) String() string {
	if !strings.HasPrefix(e.Issue, e.Field) {
		return fmt.Sprintf("%s %s", e.Field, e.Issue)
	}
	return e.Issue
}

// Errors encapsulates a list of validation errors. It also supports the
// Error interface, so it can itself be used to return an error.
type Errors []Error

// String returns a text representation of all validation errors.
func (ve Errors) String() string {
	if len(ve) == 0 {
		return ""
	}
	return ve.Join("\n")
}

// Join returns a textual representation of all validation errors,
// concatenated with the given separator.
func (ve Errors) Join(sep string) string {
	return strings.Join(ve.Errors(), sep)
}

// Errors returns an array with all errors in textual form.
func (ve Errors) Errors() []string {
	if len(ve) == 0 {
		return nil
	}
	var errors []string
	for _, e := range ve {
		errors = append(errors, e.String())
	}
	return errors
}

// Error ensures that Errors satisfies Go's error type.
func (ve Errors) Error() string {
	switch len(ve) {
	case 0:
		return ""
	case 1:
		return "Invalid record with 1 error"
	default:
		return fmt.Sprintf("Invalid record with %d errors", len(ve))
	}
}

// Add adds a validation error to the list of errors.
func (ve Errors) Add(field, issue string) Errors {
	ve = append(ve, Error{Field: field, Issue: issue})
	return ve
}

// HasErrors returns true if there are any validation errors.
func (ve Errors) HasErrors() bool {
	return len(ve) != 0
}
