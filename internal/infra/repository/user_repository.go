package repository

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
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

// Create creates a new user in the database.
//
// ctx - The context object.
// user - The user object to be created.
// Returns the key of the created user and an error, if any.
func (ur *UserRepository) Create(ctx context.Context, user *entity.User) (string, error) {

	meta, err := ur.collectionInstance.InsertDocument(user)
	if err != nil {
		return "", errors.New("failed to create user: " + err.Error())
	}
	return meta.Key, nil
}

const delUserId = "for user in User filter user.id == @id remove user in User"

func (ur *UserRepository) DeleteUserByID(ctx context.Context, id string) error {
	bindVars := map[string]interface{}{
		"id": id,
	}
	cr, err := ur.DB.Database.Query(ctx, delUserId, bindVars)
	if err != nil {
		return errors.New("failed to delete user: " + err.Error())
	}
	defer cr.Close()

	return nil
}

const getUserId = "for user in User filter user.id == @id return user"

func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	bindeVars := map[string]interface{}{
		"id": id,
	}
	cr, err := ur.DB.Database.Query(ctx, getUserId, bindeVars)
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

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	bindeVars := map[string]interface{}{
		"email": email,
	}
	cr, err := ur.DB.Database.Query(ctx, getUserEmail, bindeVars)
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

func (ur *UserRepository) ListAll(ctx context.Context, limit int) ([]*entity.User, error) {
	var out []*entity.User
	bindVars := map[string]interface{}{
		"limit": limit,
	}
	cr, err := ur.DB.Database.Query(ctx, listAllUser, bindVars)
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

func (ur *UserRepository) UpdateUserByID(ctx context.Context, id string, user *entity.User) error {
	bindVars := map[string]interface{}{
		"id":   id,
		"user": user,
	}
	cr, err := ur.DB.Database.Query(ctx, updateUserId, bindVars)
	if err != nil {
		return errors.New("failed to update user: " + err.Error())
	}
	defer cr.Close()
	var updatedUser *entity.User
	_, err = cr.ReadDocument(ctx, &updatedUser)
	if err != nil {
		if driver.IsNoMoreDocuments(err) {
			return errors.New("user not found with id: " + id)
		}
		return errors.New("failed to update user: " + err.Error())
	}

	return nil
}

const getKeyFromUserId = "for user in User filter user.id == @id return user._key"

func (ur *UserRepository) GetKeyFromUserId(ctx context.Context, userID string) (*string, error) {
	bindVars := map[string]interface{}{
		"id": userID,
	}
	cr, err := ur.DB.Database.Query(ctx, getKeyFromUserId, bindVars)
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
