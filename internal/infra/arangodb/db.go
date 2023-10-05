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

// SelectQuery is a function that executes a select AQL query on the database.
//
// It takes a query string and a slice of interface{} as its parameters.
// The query string represents the select query to be executed.
// The results slice is used to store the query results.
//
// The function returns an error if the query execution or result reading fails.
// Otherwise, it returns nil.
func (db *DBCollection) SelectQuery(query string, results interface{}) error {
	cursor, err := db.collection.Database().Query(db.ctx, query, nil)
	out := make([]interface{}, 0)
	if err != nil {
		return errors.New("Failed to query: " + err.Error())
	}
	defer cursor.Close()
	for cursor.HasMore() {
		var result interface{}
		_, err := cursor.ReadDocument(db.ctx, result)
		if err != nil {
			return errors.New("Failed to read document: " + err.Error())
		}
		out = append(out, result)
	}
	results = out
	return nil
}
