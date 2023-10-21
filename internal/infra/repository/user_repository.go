package repository

import (
	"context"
	"errors"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
	"github.com/gofiber/fiber/v2/log"
)

type UserRepository struct {
	arangodb       *arangodb.DB
	collectionName string
	dbCollection   *arangodb.DBCollection
}

func NewUserRepository(db *arangodb.DB, collectionName string) *UserRepository {
	coll := db.FromCollection(collectionName)
	return &UserRepository{
		arangodb:       db,
		collectionName: collectionName,
		dbCollection:   coll,
	}
}

type UserDocument struct {
	Key       string `json:"_key"`
	Id        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Avatar    []byte `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewUserDocument(user *entity.User) *UserDocument {
	return &UserDocument{
		Key:       user.ID.String(),
		Id:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

// Create creates a new user in the database.
//
// ctx - The context object.
// user - The user object to be created.
// Returns the key of the created user and an error, if any.
func (ur *UserRepository) Create(ctx context.Context, user *entity.User) (string, error) {
	userDocument := NewUserDocument(user)
	_, err := ur.dbCollection.InsertDocument(userDocument)
	if err != nil {
		return "", errors.New("failed to create user: " + err.Error())
	}
	return user.ID.String(), nil
}

func (ur *UserRepository) DeleteUserByID(ctx context.Context, id string) error {
	_, err := ur.dbCollection.DeleteDocument(id)
	if err != nil {
		if driver.IsDataSourceOrDocumentNotFound(err) {
			return errors.New("user not found with id: " + id)
		}
		return errors.New("failed to delete user: " + err.Error())
	}
	return nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	var usr entity.User
	_, err := ur.dbCollection.GetByKey(id, &usr)
	if err != nil {
		if driver.IsDataSourceOrDocumentNotFound(err) {
			return nil, errors.New("user not found with id: " + id)
		}
		return nil, errors.New("failed to get user by id: " + err.Error())
	}
	return &usr, nil
}

const getUserEmail = "for user in User filter user.email == @email return user"

// GetUserByEmail retrieves a user from the UserRepository by their email.
//
// ctx: The context for the function execution.
// email: The email of the user to retrieve.
// Returns a pointer to the User entity and an error, if any.
func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	bindeVars := map[string]interface{}{
		"email": email,
	}
	cr, err := ur.arangodb.Database.Query(ctx, getUserEmail, bindeVars)
	if err != nil {
		return nil, errors.New("failed to get user by email: " + err.Error())
	}
	defer cr.Close()

	var usr *entity.User
	_, err = cr.ReadDocument(ctx, &usr)
	if err != nil {
		if driver.IsNoMoreDocuments(err) {
			return nil, errors.New("user not found with email: " + email)
		}
		return nil, errors.New("failed to get user by email: " + err.Error())
	}

	return usr, nil
}

const listAllUser = "for user in User limit @limit return user"

// ListAll retrieves all users from the UserRepository.
//
// It takes two parameters:
// 1. ctx: the context.Context object for cancellation and timeout.
// 2. limit: the maximum number of users to retrieve.
//
// It returns a slice of *entity.User and an error.
func (ur *UserRepository) ListAll(ctx context.Context, limit int) ([]*entity.User, error) {
	var out []*entity.User
	bindVars := map[string]interface{}{
		"limit": limit,
	}
	cr, err := ur.arangodb.Database.Query(ctx, listAllUser, bindVars)
	if err != nil {
		return nil, errors.New("failed to list all users: " + err.Error())
	}
	defer cr.Close()
	for cr.HasMore() {
		var usr *entity.User
		_, err := cr.ReadDocument(ctx, &usr)
		if err != nil {
			return nil, errors.New("failed to retrieve all users: " + err.Error())
		}
		out = append(out, usr)
	}
	return out, nil
}

func (ur *UserRepository) UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error) {
	var updatedUser *entity.User
	log.Debug("Update: " + id)
	_, err := ur.Collection().UpdateDocument(driver.WithReturnNew(ctx, &updatedUser), id, user)
	if err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}
	return updatedUser, nil
}

// Colletion returns the collection instance of the UserRepository.
//
// No parameters.
// Returns *driver.Collection.
func (ur *UserRepository) Collection() driver.Collection {
	return *ur.dbCollection.Collection()
}

// Database returns the database instance of the UserRepository.
//
// No parameters.
// Returns a pointer to the driver.Database.
func (ur *UserRepository) Database() *driver.Database {
	return &ur.arangodb.Database
}
