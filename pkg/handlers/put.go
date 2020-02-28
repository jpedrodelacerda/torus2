package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

func (s *service) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	s.userRepository.Log("Handle PUT Request at", r.URL.Path)

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
	user := nodb.User{}

	s.userRepository.Log("trying to update user")
	user.FromJSON(r.Body)

	err = s.userRepository.UpdateUserByIndex(id, user)
	if err != nil {
		ne := NewQueryError(err.Error(), s.Docs())
		rw.WriteHeader(http.StatusInternalServerError)
		ne.ToJSON(rw)
		return
	}
	s.userRepository.Log("user successfully updated")
}
