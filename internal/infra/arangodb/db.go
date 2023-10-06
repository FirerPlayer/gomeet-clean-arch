package arangodb

import (
	"context"
	"errors"

	driver "github.com/arangodb/go-driver"
)

type DBCollection struct {
	ctx        context.Context
	collection driver.Collection
}

type DB struct {
	ctx         context.Context
	client      driver.Client
	Database    driver.Database
	Collections map[string]*DBCollection
}

func newDB(ctx context.Context, client driver.Client, database driver.Database, mapCollections map[string]*DBCollection) *DB {
	return &DB{
		ctx:         ctx,
		client:      client,
		Database:    database,
		Collections: mapCollections,
	}
}

// FromCollection returns the DBCollection associated with the given collection name.
//
// Parameters:
// - collectionName: the name of the collection.
//
// Returns:
// - *DBCollection: the DBCollection associated with the given collection name.
func (db *DB) FromCollection(collectionName string) *DBCollection {
	return db.Collections[collectionName]
}

// InsertDocument inserts a document into the DBCollection.
//
// It takes a single parameter, v, which is the document to be inserted.
// The function returns a pointer to a DocumentMeta and an error.
func (db *DBCollection) InsertDocument(v interface{}) (*driver.DocumentMeta, error) {
	document, err := db.collection.CreateDocument(db.ctx, v)
	if err != nil {
		return nil, errors.New("Failed to create document: " + err.Error())
	}
	return &document, nil
}

// DeleteDocument deletes a document from the DBCollection.
//
// Parameters:
// - key: the key of the document to be deleted.
//
// Returns:
// - *driver.DocumentMeta: the metadata of the deleted document.
// - error: an error if the deletion fails.
func (db *DBCollection) DeleteDocument(key string) (*driver.DocumentMeta, error) {
	meta, err := db.collection.RemoveDocument(db.ctx, key)
	if err != nil {
		return nil, errors.New("Failed to delete document: " + err.Error())
	}
	return &meta, nil
}

// GetByKey retrieves a document from the database by its key.
//
// Parameters:
// - key: the key of the document to retrieve.
// - result: a pointer to a variable to store the retrieved document.
//
// Returns:
// - *driver.DocumentMeta: metadata of the retrieved document.
// - error: any error that occurred during the retrieval process.
func (db *DBCollection) GetByKey(key string, result interface{}) (*driver.DocumentMeta, error) {
	meta, err := db.collection.ReadDocument(db.ctx, key, result)
	if err != nil {
		return nil, errors.New("Failed to read document: " + err.Error())
	}
	return &meta, nil

}

// ExecQuery is a function that executes a query on a database collection.
//
// It takes a query string and a map of bind variables as parameters.
// The query string is the query to be executed on the database.
// The bindVars map contains the bind variables to be used in the query.
// It returns an error if the query execution fails.
func (db *DBCollection) ExecQuery(query string, bindVars map[string]interface{}) error {
	cr, err := db.collection.Database().Query(db.ctx, query, bindVars)
	if err != nil {
		return errors.New("Failed to exec query: " + err.Error())
	}
	defer cr.Close()
	return nil
}
