package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ehehe90/gqlgen-study/graph"
	// "github.com/ehehe90/gqlgen-study/graph/model"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"
const dataSource = "root:password@tcp(127.0.0.1:3306)/gqlgen_study?charset=utf8mb4&parseTime=true"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic (err)
			}
		}
	}()
	db.LogMode(true)
	// db.AutoMigrate(&model.User{}, &model.Todo{})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
