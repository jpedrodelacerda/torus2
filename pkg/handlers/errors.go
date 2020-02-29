package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	ErrParsingURL            = errors.New("error while parsing URL")
	ErrDeserializationFailed = errors.New("failed to deserialize data")
)

// QueryError encapsulates the given error and the documentation URL from the API
// swagger:model QueryError
type QueryError struct {
	// error message
	Message string `json:"message"`
	// documentation link
	DocumentationURL string `json:"documentation_url"`
}

// NewQueryError is the idiomatic way of creating a new QueryError struct
func NewQueryError(msg string, doc string) QueryError {
	return QueryError{
		Message:          msg,
		DocumentationURL: doc,
	}
}

// Unwrap returns the error message from the
func (q *QueryError) Unwrap() string {
	return q.Message
}

// Error method returns the error string
func (q *QueryError) Error() string {
	return fmt.Sprint(q.Message)
}

// ToJSON marshalls the query error struct
func (q *QueryError) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(q)
}
