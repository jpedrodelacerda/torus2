package nodb

// swagger:model
// User struct used by the database
type User struct {
	// the id for the user
	//
	// required: false
	// min: 1
	ID int `json:"id"`
	// the email for the user
	//
	// required: true
	// example: user@example.com
	Email string `json:"email" validate:"required,email"`
	// the password hash for the user
	//
	// min length: 1
	PasswordHash string `json:"password_hash"`
}
