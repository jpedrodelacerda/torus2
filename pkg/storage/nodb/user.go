package nodb

// User struct used by the database
// swagger:model User
type User struct {

	// the id for the user
	//
	// required: false
	// min: 1
	ID int `json:"id"`

	// the email for the user
	//
	// required: true
	// max length: 255
	// swagger:strfmt email
	Email string `json:"email" validate:"required,email"`

	// the password hash for the user
	//
	// min length: 1
	PasswordHash string `json:"password_hash"`
}

// NOTE: the following types have been defined by documentation purposes
// they are not used by any handlers
