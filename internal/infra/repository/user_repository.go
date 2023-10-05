package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
)

type UserRepository struct {
	DB                 *arangodb.DB
	CollectionName     string
	collectionInstance *arangodb.DBCollection
}

func NewUserRepository(db *arangodb.DB, collectionName string) (*UserRepository, error) {
	coll := db.FromCollection(collectionName)
	return &UserRepository{
		DB:                 db,
		CollectionName:     collectionName,
		collectionInstance: coll,
	}, nil
}

// Create creates a new user in the User repository.
//
// ctx - The context object.
// user - The user object to be created.
// Returns the key of the created user and an error, if any.
func (ur *UserRepository) Create(ctx context.Context, user interface{}) (string, error) {
	meta, err := ur.collectionInstance.InsertDocument(user)
	if err != nil {
		return "", errors.New("Failed to create user: " + err.Error())
	}
	return meta.Key, nil
}

var delById = func(q string) string {
	return fmt.Sprintf(
		`for user in User 
			filter user.id == %s 
			remove user in User`,
		q,
	)
}

func (ur *UserRepository) DeleteUserByID(ctx context.Context, id string) error {
	cr, err := ur.DB.Database.Query(ctx, delById(id), nil)
	if err != nil {
		return errors.New("Failed to delete user: " + err.Error())
	}
	err = cr.Close()
	if err != nil {
		return errors.New("Failed to delete user: " + err.Error())
	}

	return nil
}
func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return nil, nil
}
func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}
func (ur *UserRepository) ListAll(ctx context.Context, limit int) ([]*entity.User, error) {
	return nil, nil
}
func (ur *UserRepository) UpdateUserByID(ctx context.Context, id string, user *entity.User) error {
	return nil
}
