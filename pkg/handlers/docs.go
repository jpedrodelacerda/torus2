// Package classification Torus API
//
// The purpose of this application is to provide means to manage Torus system
//
// Terms Of Service:
// None. Use at your own risk.
//
//		Schemes: http
// 		BasePath: /
// 		Version: 0.0.1
//	    License: MIT http://opensource.org/licenses/MIT
//		Contact: Torus <tech@torus.com> https://torus.com
//
// 		Consumes:
// 		- application/json
//
// 		Produces:
// 		- application/json
// swagger:meta

package handlers

import "github.com/jpedrodelacerda/torus2/pkg/storage/nodb"

// swagger:parameters userID
// The ID from the specified user
type userIDParameterWrapper struct {
	// The user id code
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response userListResponse
// A list of users
type userListResponseWrapper struct {
	// All current products
	// in: body
	Body nodb.UserList
}

// swagger:response userResponse
// Data structure representing a single user
type userResponseWrapper struct {
	// Newly created user
	// in: body
	Body nodb.User
}

// swagger:response errorResponse
// structured error returned as a json
type errorResponseWrapper struct {
	// User not found message
	// in: body
	Body QueryError
}
