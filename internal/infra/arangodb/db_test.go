package arangodb

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	driver "github.com/arangodb/go-driver"
// 	"github.com/arangodb/go-driver/http"
// 	"github.com/matryer/is"
// )

// type mockDoc struct {
// 	name string
// 	date time.Time
// }

// func TestDB(t *testing.T) {
// 	// GetByKey := gomega.NewWithT(t)
// 	is := is.New(t)
// 	ctx := context.Background()

// 	mockConnection, err := http.NewConnection(http.ConnectionConfig{
// 		Endpoints: []string{"http://localhost:8529"},
// 	})
// 	is.NoErr(err) // must have connection

// 	mockClient, err := driver.NewClient(driver.ClientConfig{
// 		Connection:     mockConnection,
// 		Authentication: driver.BasicAuthentication("root", "root"),
// 	})
// 	is.NoErr(err) // must have client

// 	mockDb, err := mockClient.CreateDatabase(ctx, "test", nil)
// 	is.NoErr(err) // must have database
// 	mockCollection, err := mockDb.CreateCollection(ctx, "test_collection", nil)
// 	is.NoErr(err) // must create a collection
// 	mockMeta, err := mockCollection.CreateDocument(ctx, mockDoc{name: "test", date: time.Now()})
// 	is.NoErr(err) // must create a document

// 	db, err := NewDB(ctx, mockClient, "test")
// 	is.NoErr(err) // need to found the db test
// 	coll, err := db.FromCollection("test_collection")
// 	is.NoErr(err)                                              // need to found the collection test_collection
// 	is.Equal(coll.collection.Name(), mockCollection.Name())    // collections must be the same
// 	is.Equal(coll.collection.Database().Name(), mockDb.Name()) // databases must be the same

// 	t.Run("Collection Insert", func(t *testing.T) {
// 		is := is.New(t)

// 		metaKey, err := coll.InsertDocument(mockDoc{name: "test", date: time.Now()})
// 		is.NoErr(err)                        // must create a document
// 		is.True(metaKey.Key != mockMeta.Key) // key must be different

// 		var docs = make([]mockDoc, 2)
// 		_, an, err := mockCollection.ReadDocuments(ctx, []string{metaKey.Key, mockMeta.Key}, docs)
// 		fmt.Println(an)
// 		is.NoErr(err)          // must read documents
// 		is.Equal(len(docs), 2) // must have 2 documents
// 	})
// }
