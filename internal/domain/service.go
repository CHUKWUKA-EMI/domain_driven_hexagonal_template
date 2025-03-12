package domain

// Service interface
type Service interface {
	// FindUserByID returns a user by ID
	FindUserByID(id string) (*User, error)
	// FindAllUsers returns all users
	FindAllUsers() ([]*User, error)
	// CreateUser creates a new user
	CreateUser(firstName, lastName, email, street, city, state, zipCode string) (*User, error)
	// UpdateUser updates a user
	UpdateUser(id, firstName, lastName, email, street, city, state, zipCode string) (*User, error)
	// DeleteUser deletes a user by ID
	DeleteUser(id string) error
}
