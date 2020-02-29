package handlers

import (
	"errors"
	"net/http"

	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

// swagger:route POST /users/{id} users registerUser
// Registers a new user
//
// Responses:
// 	201: userResponse
//  422: errorResponse
// 	500: errorResponse

// AddUser handles POST requests to register new users on the database
func (s *service) AddUser(rw http.ResponseWriter, r *http.Request) {
	s.userRepository.Log("Handle POST Request at ", r.URL.Path)

	user := r.Context().Value(contextKey("user")).(nodb.User)

	err := s.userRepository.AddUser(user)
	if errors.Is(err, nodb.ErrExistingUser) {
		s.userRepository.Log("user already exists")
		ne := NewQueryError(err.Error(), "localhost:8080/docs")
		rw.WriteHeader(http.StatusBadRequest)
		ne.ToJSON(rw)
		return
	}
	if err != nil {
		s.userRepository.Log("Failed to add user")
		http.Error(rw, "Error adding user to database", http.StatusInternalServerError)
		return
	}

	user.ToJSON(rw)
	rw.WriteHeader(http.StatusCreated)
}
