package main

import (
	"github.com/firerplayer/whatsmeet-go/internal/infra/web/webserver"
)

func main() {

	server := webserver.NewWebServer("8080", "GoMeet")
	err := server.Start()
	if err != nil {
		panic(err)
	}

	// ctx := context.Background()

	// mockConnection, err := http.NewConnection(http.ConnectionConfig{
	// 	Endpoints: []string{"http://localhost:8529"},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// client, err := driver.NewClient(driver.ClientConfig{
	// 	Connection:     mockConnection,
	// 	Authentication: driver.BasicAuthentication("root", "root"),
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// db, err := arangodb.NewDB(ctx, client, "gommet")
	// if err != nil {
	// 	panic(err)
	// }
	// coll, err := db.FromCollection("denis")
	// if err != nil {
	// 	panic(err)
	// }

	// mockDb, err := mockClient.CreateDatabase(ctx, "test", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// mockCollection, err := mockDb.CreateCollection(ctx, "test_collection", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// mockMeta, err := coll.InsertDocument(entity.NewUser("#$!@ Denis 222", "denis@email.com 44", "This  asasasd is my bio", nil))

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(mockMeta)

}
