package arangodb

import (
	"context"
	"errors"
	"log"

	driver "github.com/arangodb/go-driver"
)

type DB struct {
	client         driver.Client
	databaseName   string
	collectionName string
	collection     driver.Collection
	database       driver.Database
	ctx            context.Context
}

func NewDB(client driver.Client, database string) (*DB, error) {
	newDb := &DB{
		client:       client,
		databaseName: database,
	}
	err := newDb.setup()
	if err != nil {
		return nil, errors.New("Failed to setup database: " + err.Error())
	}
	return newDb, nil
}

func (db *DB) setup() error {
	var err error
	db.database, err = db.client.Database(db.ctx, db.databaseName)
	if err != nil {
		return err
	}
	db.collection, err = db.database.Collection(db.ctx, db.collectionName)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) InsertDocument(v interface{}) string {
	document, err := db.collection.CreateDocument(db.ctx, v)
	if err != nil {
		log.Fatalf("Failed to create document: %v", err)
	}
	return document.Key
}

func (db *DB) DeleteDocument(key string) (*driver.DocumentMeta, error) {
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
func (db *DB) GetByKey(key string, result interface{}) (*driver.DocumentMeta, error) {
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
func (db *DB) SelectQuery(query string, results []interface{}) error {
	cursor, err := db.collection.Database().Query(db.ctx, query, nil)
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
		results = append(results, result)
	}
	return nil
}
