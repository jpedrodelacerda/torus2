// Package classification Torus API
//
// The purpose of this application is to provide means to manage Torus system
//
// Terms Of Service:
// None. Use at your own risk.
//
//	Schemes: http
//	BasePath: /
//	Version: 0.0.1
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: Torus <tech@torus.com> https://torus.com
//
//	Consumes:
//	- application/json
//
// 	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/jpedrodelacerda/torus2/pkg/storage/nodb"

// The ID from the specified user
// swagger:parameters fetchUser deleteUser updateUser
type userIDParameterWrapper struct {
	// The user id code
	// in: path
	// required: true
	ID int `json:"id"`
}

// A list of users
// swagger:response userListResponse
type userListResponseWrapper struct {
	// All current users
	// in: body
	Body []nodb.User
}

// Data structure representing a single user
// swagger:response userResponse
type userResponseWrapper struct {
	// Newly created user
	// in: body
	Body nodb.User
}

// swagger:parameters registerUser updateUser
type userParamsWrapper struct {
	// User structure to be created or updated
	//
	// NOTE: The id parameter is ignored by create and update actions.
	// in: body
	Body nodb.User
}

// Structured error returned as a json
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Error structure with error and link to documentation
	// in: body
	Body QueryError
}

// no reply from server
// swagger:response noContent
type noContentResponseWrapper struct{}
