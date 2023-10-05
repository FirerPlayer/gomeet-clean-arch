package main

import (
	"context"
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/firerplayer/whatsmeet-go/internal/domain/entity"
	"github.com/firerplayer/whatsmeet-go/internal/infra/arangodb"
	"github.com/firerplayer/whatsmeet-go/internal/infra/repository"
)

func main() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		panic(err)
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "root"),
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	initializer := arangodb.NewDBInitializer(ctx, client)

	db, err := initializer.Init("gommet")
	if err != nil {
		panic(err)
	}

	_, err = repository.NewUserRepository(db, "User")
	if err != nil {
		panic(err)
	}
	var keys = make([]string, 0)

	for i := 0; i < 20; i++ {
		aux := struct {
			_key string
			user *entity.User
		}{
			_key: fmt.Sprint(i),
			user: entity.NewUser("Denis "+fmt.Sprint(i), "denis"+fmt.Sprint(i)+"@email.com", "This is my bio "+fmt.Sprint(i), nil),
		}
		dd, err := client.Database(ctx, "gommet")
		if err != nil {
			panic(err)
		}
		clll, err := dd.Collection(ctx, "User")
		if err != nil {
			panic(err)
		}
		meta, err := clll.CreateDocument(ctx, aux)
		if err != nil {
			panic(err)
		}
		keys = append(keys, meta.Key)

		// meta, err := db.FromCollection("User").InsertDocument(entity.NewUser("Denis "+fmt.Sprint(i), "denis"+fmt.Sprint(i)+"@email.com", "This is my bio "+fmt.Sprint(i), nil))
		// if err != nil {
		// 	panic(err)
		// }
		// keys = append(keys, meta.Key)
	}
	fmt.Println(keys)
	// err = userRepo.DeleteUserByID(ctx, keys[len(keys)-1])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User deleted at key: ", keys[len(keys)-1])
	// db.FromCollection("User").InsertDocument(entity.NewUser("Denis222", "denis222@email.com", "This is my bio 222", nil))

	// server := webserver.NewWebServer("8080", "GoMeet")

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

	// userRepository, err := repository.NewUserRepository(db, "User")
	// if err != nil {
	// 	panic(err)
	// }

	// createUserUsecase := user.NewCreateUserUsecase(userRepository)
	// server.Get("/", func(c *fiber.Ctx) error {
	// 	user, err := createUserUsecase.Execute(c.Context(), dto.CreateUserInputDTO{
	// 		Name:  "Denis",
	// 		Email: "denis@email.com",
	// 		Bio:   "This is my bio",
	// 	})
	// 	if err != nil {
	// 		c.SendStatus(500)
	// 		return c.SendString("Failed to create user: " + err.Error())
	// 	}

	// 	return c.JSON(user)
	// })

	// err = server.Start()
	// if err != nil {
	// 	panic(err)
	// }

	// ctx := context.Background()

	// mockConnection, err := http.NewConnection(http.ConnectionConfig{
	// 	Endpoints: []string{"http://localhost:8529"},
	// })
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
	//driver

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
