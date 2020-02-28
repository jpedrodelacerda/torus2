package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jpedrodelacerda/torus2/pkg/storage/nodb"
)

// Service provides the REST capabilities
type Service interface {
	ListUsers(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	FetchUser(http.ResponseWriter, *http.Request)
	AddUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
	MiddlewareValidateUser(http.Handler) http.Handler
	Addr() string
	Docs() string
}

// Repository provides interaction to the database
type Repository interface {
	AddUser(nodb.User) error
	DeleteUserByID(int) error
	UpdateUserByIndex(int, nodb.User) error
	FindUserByID(int) (int, nodb.User, error)
	GetUsers() nodb.UserList
	Log(v ...interface{})
}

type service struct {
	addr           string
	port           int
	userRepository Repository
}

// NewService is the idiomatic way of creating a service
func NewService(addr string, port int, repo Repository) Service {
	return &service{addr, port, repo}
}

// Host returns the hostname string
func (s *service) Addr() string {
	return fmt.Sprintf("%s:%d", s.addr, s.port)
}

func (s *service) Docs() string {
	return fmt.Sprintf("%s/docs", s.Addr())
}

func getUserID(r *http.Request) int {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}
