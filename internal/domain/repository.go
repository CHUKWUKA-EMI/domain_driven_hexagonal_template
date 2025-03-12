package domain

// Repository interface
type Repository interface {
	// FindByID returns a user by ID
	FindByID(id string) (*User, error)
	// FindAll returns all users
	FindAll() ([]*User, error)
	// Save saves a user
	Save(user *User) (*User, error)
	// Update updates a user
	Update(user *User) (*User, error)
	// Delete deletes a user by ID
	Delete(id string) error
}
