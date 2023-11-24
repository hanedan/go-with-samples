package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	d "go-with-samples/pkg/db"
	u "go-with-samples/pkg/db/user"
	mu "go-with-samples/pkg/model/user"
	s "go-with-samples/pkg/setup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
	defer cancel()

	db, err := d.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = s.CreateUserTable(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	// dependency injection
	userDB := u.NewUserDB(db)
	userApi := mu.NewUserAPI(userDB)

	router := gin.Default()
	// If you manage a single service under one domain,
	// the endpoint would be: your-service.com/v1/library-api/
	v1 := router.Group("/v1")
	// If you manage multiple services under one domain,
	// the endpoint would be: your-domain.com/library-api/v1/

	api := v1.Group("/library-api")
	books := api.Group("/users")
	books.POST("/create", userApi.CreateHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
