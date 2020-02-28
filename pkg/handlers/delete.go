package handlers

import (
	"net/http"

	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

func (s *service) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	s.userRepository.Log("Handle DELETE Request at", r.URL.Path)
	id := getUserID(r)

	err := s.userRepository.DeleteUserByID(id)
	if err == nodb.ErrUserNotFound {
		err := NewQueryError(err.Error(), s.Docs())
		err.ToJSON(rw)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		err := NewQueryError("[ERROR] unable to delete record", s.Docs())
		err.ToJSON(rw)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
