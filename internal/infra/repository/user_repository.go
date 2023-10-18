package repository

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
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

// Create creates a new user in the database.
//
// ctx - The context object.
// user - The user object to be created.
// Returns the key of the created user and an error, if any.
func (ur *UserRepository) Create(ctx context.Context, user *entity.User) (string, error) {

	_, err := ur.dbCollection.InsertDocument(user)
	if err != nil {
		return "", errors.New("failed to create user: " + err.Error())
	}
	return user.ID.String(), nil
}

const delUserId = `FOR user IN User FILTER user.id == @id REMOVE user IN User`

// DeleteUserByID deletes a user by ID.
//
// ctx: the context.Context object for cancellation and timeout.
// id: the ID of the user to be deleted.
// error: an error if the deletion fails.
func (ur *UserRepository) DeleteUserByID(ctx context.Context, id string) error {
	bindVars := map[string]interface{}{
		"id": id,
	}
	cr, err := ur.arangodb.Database.Query(ctx, delUserId, bindVars)
	if err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}
	defer cr.Close()

	return nil
}

const getUserId = "for user in User filter user.id == @id return user"

// GetUserByID retrieves a user from the UserRepository by their ID.
//
// It takes the following parameters:
// - ctx: the context.Context used for the operation.
// - id: the ID of the user to retrieve.
//
// It returns a pointer to entity.User and an error.
func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	bindeVars := map[string]interface{}{
		"id": id,
	}
	cr, err := ur.arangodb.Database.Query(ctx, getUserId, bindeVars)
	if err != nil {
		return nil, errors.New("failed to get user by id: " + err.Error())
	}
	defer cr.Close()
	var usr *entity.User
	_, err = cr.ReadDocument(ctx, &usr)
	if err != nil {
		if driver.IsNoMoreDocuments(err) {
			return nil, errors.New("user not found with id: " + id)
		}
		return nil, errors.New("failed to get user by email: " + err.Error())
	}

	return usr, nil
}

const getUserEmail = "for user in User filter user.Email == @email return user"

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

const updateUserId = "for u in User filter u.id == @id update u with @user in User return NEW"

// UpdateUserByID updates a user in the UserRepository by ID.
//
// It takes the following parameters:
// - ctx: the context.Context object for the request.
// - id: the ID of the user to update.
// - user: the updated user entity.
//
// It returns:
// - *entity.User: the updated user entity if the update is successful.
// - error: an error if the update fails.
func (ur *UserRepository) UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error) {
	bindVars := map[string]interface{}{
		"id":   id,
		"user": user,
	}
	cr, err := ur.arangodb.Database.Query(ctx, updateUserId, bindVars)
	if err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}
	defer cr.Close()
	var updatedUser *entity.User
	_, err = cr.ReadDocument(ctx, &updatedUser)
	if err != nil {
		if driver.IsNoMoreDocuments(err) {
			return nil, errors.New("user not found with id: " + id)
		}
		return nil, errors.New("failed to update user: " + err.Error())
	}

	return updatedUser, nil
}

const getKeyFromUserId = "for user in User filter user.id == @id return user._key"

// GetKeyFromUserId retrieves the key associated with a given user ID.
//
// ctx: The context.Context object for cancellation and timeouts.
// userID: The ID of the user.
// *string: The key associated with the user ID, or nil if not found.
// error: An error if there was a problem retrieving the key.
func (ur *UserRepository) GetKeyFromUserId(ctx context.Context, userID string) (*string, error) {
	bindVars := map[string]interface{}{
		"id": userID,
	}
	cr, err := ur.arangodb.Database.Query(ctx, getKeyFromUserId, bindVars)
	if err != nil {
		return nil, errors.New("failed to get key from user id: " + err.Error())
	}
	defer cr.Close()
	var key *string
	_, err = cr.ReadDocument(ctx, &key)
	if err != nil {
		if driver.IsNoMoreDocuments(err) {
			return nil, errors.New("user not found with id: " + userID)
		}
		return nil, errors.New("failed to get key from user id: " + err.Error())
	}

	return key, nil
}

// Colletion returns the collection instance of the UserRepository.
//
// No parameters.
// Returns *driver.Collection.
func (ur *UserRepository) Colletion() *driver.Collection {
	return ur.dbCollection.Collection()
}

// Database returns the database instance of the UserRepository.
//
// No parameters.
// Returns a pointer to the driver.Database.
func (ur *UserRepository) Database() *driver.Database {
	return &ur.arangodb.Database
}
