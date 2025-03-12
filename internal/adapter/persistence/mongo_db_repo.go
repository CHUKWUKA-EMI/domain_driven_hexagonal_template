package persistence

import (
	"backend_api_template/internal/domain"
	"backend_api_template/internal/infrastructure/constants"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBUserRepository is a struct that implements the domain.Repository interface
type MongoDBUserRepository struct {
	collection *mongo.Collection
}

// NewMongoDBUserRepository creates a new instance of MongoDBUserRepository
func NewMongoDBUserRepository(db *mongo.Database) domain.Repository {
	return &MongoDBUserRepository{collection: db.Collection(constants.UsersCollection)}
}

// FindAll returns all users
func (r *MongoDBUserRepository) FindAll() ([]*domain.User, error) {
	panic("implement me")
}

// FindByID returns a user by id
func (r *MongoDBUserRepository) FindByID(id string) (*domain.User, error) {
	panic("implement me")
}

// Save creates a new user
func (r *MongoDBUserRepository) Save(user *domain.User) (*domain.User, error) {
	panic("implement me")
}

// Update updates a user
func (r *MongoDBUserRepository) Update(user *domain.User) (*domain.User, error) {
	panic("implement me")
}

// Delete deletes a user
func (r *MongoDBUserRepository) Delete(id string) error {
	panic("implement me")
}
