package application

import "backend_api_template/internal/domain"

// UserServiceImpl is a struct that implements the domain.Service interface
type UserServiceImpl struct {
	repository domain.Repository
}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService(repository domain.Repository) domain.Service {
	return &UserServiceImpl{repository: repository}
}

// CreateUser implements domain.Service.
func (u *UserServiceImpl) CreateUser(firstName string, lastName string, email string, street string, city string, state string, zipCode string) (*domain.User, error) {
	return u.repository.Save(domain.NewUser("", firstName, lastName, email, domain.NewAddress(street, city, state, zipCode)))
}

// DeleteUser implements domain.Service.
func (u *UserServiceImpl) DeleteUser(id string) error {
	return u.repository.Delete(id)
}

// FindAllUsers implements domain.Service.
func (u *UserServiceImpl) FindAllUsers() ([]*domain.User, error) {
	return u.repository.FindAll()
}

// FindUserByID implements domain.Service.
func (u *UserServiceImpl) FindUserByID(id string) (*domain.User, error) {
	return u.repository.FindByID(id)
}

// UpdateUser implements domain.Service.
func (u *UserServiceImpl) UpdateUser(id string, firstName string, lastName string, email string, street string, city string, state string, zipCode string) (*domain.User, error) {
	return u.repository.Update(domain.NewUser(id, firstName, lastName, email, domain.NewAddress(street, city, state, zipCode)))
}
