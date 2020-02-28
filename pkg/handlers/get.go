package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

// swagger:route GET /users users listUsers
// Returns a list of users registered at the system
// Responses:
// 	200: userListResponse
func (s *service) ListUsers(rw http.ResponseWriter, r *http.Request) {
	s.userRepository.Log("Handle GET Request at", r.URL.Path)

	ul := s.userRepository.GetUsers()

	err := ul.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /users/{id} users fetchUser
// Returns a specific user registered
// Responses:
// 	200: userResponse
// 	404: errorResponse
func (s *service) FetchUser(rw http.ResponseWriter, r *http.Request) {
	s.userRepository.Log("Handle GET Request at", r.URL.Path)

	vars := mux.Vars(r)
	if len(vars) != 1 {
		ne := NewQueryError(ErrParsingURL.Error(), s.Docs())
		rw.WriteHeader(http.StatusBadRequest)
		ne.ToJSON(rw)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ne := NewQueryError(ErrParsingURL.Error(), s.Docs())
		rw.WriteHeader(http.StatusBadRequest)
		ne.ToJSON(rw)
		return
	}

	_, user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		s.userRepository.Log("error while fetching user:", err)
		ne := NewQueryError(nodb.ErrUserNotFound.Error(), s.Docs())
		rw.WriteHeader(http.StatusNotFound)
		ne.ToJSON(rw)
		return
	}
	s.userRepository.Log("user found")
	rw.WriteHeader(http.StatusOK)
	user.ToJSON(rw)
}
