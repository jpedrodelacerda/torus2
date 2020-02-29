package handlers

import (
	"context"
	"net/http"

	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

type contextKey string

func (s *service) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s.userRepository.Log("[Middleware] [Validation] Handle validation at", r.URL.Path)
		user := nodb.User{}

		err := user.FromJSON(r.Body)
		if err != nil {
			s.userRepository.Log("[ERROR] [Middleware] [Validation] Failed to deserialize user")
			ne := NewQueryError(err.Error(), s.Docs())
			rw.WriteHeader(http.StatusUnprocessableEntity)
			ne.ToJSON(rw)
			return
		}

		err = user.Validate()
		if err != nil {
			s.userRepository.Log("[ERROR] [Middleware] [Validation] Failed to validate user")
			ne := NewQueryError(err.Error(), s.Docs())
			rw.WriteHeader(http.StatusBadRequest)
			ne.ToJSON(rw)
			return
		}

		ctx := context.WithValue(r.Context(), contextKey("user"), user)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

func (s *service) MiddlewareWriteJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	})
}
