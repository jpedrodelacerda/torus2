package nodb

import (
	"encoding/json"
	"io"
	"log"

	"github.com/go-playground/validator"
)

// Validate returns true if everything is alright
func (u *User) Validate() error {
	// return nil
	validate := validator.New()

	return validate.Struct(u)
}

// Storage is the user list itself
type Storage struct {
	l        *log.Logger
	userList UserList
}

// UserList is a collection of users
type UserList []User

// AddUser adds a new user to the list
func (s *Storage) AddUser(u User) error {
	for _, v := range s.userList {
		if u.Email == v.Email {
			return ErrExistingUser
		}
	}

	u.ID = s.nextID()
	s.l.Println("Adding user")
	s.userList = append(s.userList, u)
	return nil
}

// NewStorage is the idiomatic way to create a new storage
func NewStorage(logger *log.Logger) *Storage {
	return &Storage{userList: ul, l: logger}
}

// Log method for the storage
func (s *Storage) Log(v ...interface{}) {
	s.l.Println(v...)
}

// GetUsers lists all users on db
func (s *Storage) GetUsers() UserList {
	return s.userList
}

// ToJSON encodes the list of users to the io.Writer
func (ul *UserList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ul)
}

// ToJSON encodes the user to io.Writer
func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// UpdateUserByIndex updates the user with the given id
func (s *Storage) UpdateUserByIndex(id int, u User) error {
	pos, _, err := s.FindUserByID(id)
	if err != nil {
		return err
	}

	s.userList[pos] = u
	s.userList[pos].ID = id
	return nil
}

// FromJSON decodes the given user to the io.Reader
func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

// FindUserByID returns the index, user struct and error
func (s *Storage) FindUserByID(id int) (pos int, user User, err error) {
	for k, v := range s.userList {
		if id == v.ID {
			return k, v, nil
		}
	}
	return -1, User{}, ErrUserNotFound
}

func (s *Storage) findUserByMail(mail string) (pos int, user User, err error) {
	for k, v := range s.userList {
		if mail == v.Email {
			return k, v, nil
		}
	}
	return -1, User{}, ErrUserNotFound
}

func (s *Storage) DeleteUserByID(id int) error {
	pos, _, err := s.FindUserByID(id)
	if err != nil {
		return err
	}
	npos := pos + 1
	s.userList = append(s.userList[:pos], s.userList[npos:]...)
	return nil
}

func (s *Storage) nextID() int {
	return len(s.userList) + 1
}

var ul = []User{
	User{
		ID:           1,
		Email:        "jpdl",
		PasswordHash: "ashdua"},
	User{
		ID:           2,
		Email:        "debra",
		PasswordHash: "oioioi"},
}
