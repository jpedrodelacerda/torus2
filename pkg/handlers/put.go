package handlers

import (
	"net/http"

	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

// swagger:route PUT /users/{id} users updateUser
// Updates the user registry
// Responses:
// 	200: userResponse
//	422: errorResponse
// 	500: errorResponse

// UpdateUser handles PUT requests to update user data
func (s *service) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	s.userRepository.Log("Handle PUT Request at", r.URL.Path)

	id := getUserID(r)
	user := nodb.User{}

	s.userRepository.Log("trying to update user")
	user.FromJSON(r.Body)

	err := s.userRepository.UpdateUserByIndex(id, user)
	if err != nil {
		ne := NewQueryError(err.Error(), s.Docs())
		rw.WriteHeader(http.StatusInternalServerError)
		ne.ToJSON(rw)
		return
	}
	s.userRepository.Log("user successfully updated")
}
