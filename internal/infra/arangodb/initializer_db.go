package arangodb

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
)

var collectionNames = [3]string{
	"User",
	"Chat",
	"Message",
}

type DBInitializer struct {
	ctx    context.Context
	client driver.Client
}

func NewDBInitializer(ctx context.Context, client driver.Client) *DBInitializer {
	return &DBInitializer{
		ctx:    ctx,
		client: client,
	}
}

func (s *DBInitializer) initDatabase(dbName string) (driver.Database, error) {
	exist, err := s.client.DatabaseExists(s.ctx, dbName)
	if err != nil {
		return nil, errors.New("failed to check if database exists: " + err.Error())
	}

	if !exist {
		db, err := s.client.CreateDatabase(s.ctx, dbName, nil)
		if err != nil {
			return nil, errors.New("failed to create database: " + err.Error())
		}
		return db, nil
	}

	db, err := s.client.Database(s.ctx, dbName)
	if err != nil {
		return nil, errors.New("failed to find database: " + err.Error())
	}
	return db, nil
}

func (s *DBInitializer) initCollections(database driver.Database) (map[string]*DBCollection, error) {
	var out = make(map[string]*DBCollection)
	for _, name := range collectionNames {
		exist, err := database.CollectionExists(s.ctx, name)
		if err != nil {
			return nil, errors.New("failed to check if collection exists: " + err.Error())
		}
		// t := true
		if !exist {
			cl, err := database.CreateCollection(s.ctx, name, &driver.CreateCollectionOptions{
				KeyOptions: &driver.CollectionKeyOptions{
					AllowUserKeys: true,
				},
			})
			if err != nil {
				return nil, errors.New("failed to create collection: " + err.Error())
			}
			out[name] = &DBCollection{
				ctx:        s.ctx,
				collection: cl,
			}
		} else {
			cl, err := database.Collection(s.ctx, name)
			if err != nil {
				return nil, errors.New("failed to find collection: " + err.Error())
			}
			out[name] = &DBCollection{
				ctx:        s.ctx,
				collection: cl,
			}
		}

	}

	return out, nil
}

func (s *DBInitializer) Init(databaseName string) (*DB, error) {
	database, err := s.initDatabase(databaseName)
	if err != nil {
		return nil, errors.New("failed to init database: " + err.Error())
	}

	mapCollections, err := s.initCollections(database)
	if err != nil {
		return nil, errors.New("failed to init collections: " + err.Error())
	}

	db := newDB(s.ctx, s.client, database, mapCollections)

	return db, nil

}
