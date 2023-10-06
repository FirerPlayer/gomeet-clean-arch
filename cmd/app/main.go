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

	ur, err := repository.NewUserRepository(db, "User")
	if err != nil {
		panic(err)
	}

	users, err := ur.ListAll(ctx, 10000)
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		printUser(u)

	}

}

func printUser(u *entity.User) {
	fmt.Println("\n---------------------------------------------------------------------------")
	fmt.Println("Id: ", u.ID)
	fmt.Println("Nome: ", u.Name)
	fmt.Println("Email: ", u.Email)
	fmt.Println("Bio: ", u.Bio)
	fmt.Println("Avatar: ", u.Avatar)
	fmt.Println("CreatedAt: ", u.CreatedAt)
	fmt.Println("UpdatedAt: ", u.UpdatedAt)

}
