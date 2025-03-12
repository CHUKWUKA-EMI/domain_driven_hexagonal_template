package application

import (
	"backend_api_template/internal/domain"

	"github.com/go-playground/validator/v10"
)

// UserDTO is a struct that represents the user data transfer object
type UserDTO struct {
	ID        string `json:"id,omitempty" validate:"omitempty"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Address   struct {
		Street  string `json:"street" validate:"required"`
		City    string `json:"city" validate:"required"`
		State   string `json:"state" validate:"required"`
		ZipCode string `json:"zip_code" validate:"required"`
	} `json:"address" validate:"required"`
}

// Validate validates the UserDTO struct
func (dto *UserDTO) Validate(validateFunc *validator.Validate) error {
	return validateFunc.Struct(dto)
}

// ToDomain converts the UserDTO struct to a User domain struct
func (dto *UserDTO) ToDomain() *domain.User {
	return domain.NewUser(
		dto.ID,
		dto.FirstName,
		dto.LastName,
		dto.Email,
		domain.NewAddress(
			dto.Address.Street,
			dto.Address.City,
			dto.Address.State,
			dto.Address.ZipCode),
	)
}
