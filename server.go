package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/iandjx/go-order-graphql-api/graph"
	"github.com/iandjx/go-order-graphql-api/graph/generated"
	"github.com/iandjx/go-order-graphql-api/pkg/dbmodel"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func initDB() *gorm.DB {
	var err error
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&dbmodel.Order{}, &dbmodel.Item{})
	return db
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
