package main

import (
	"github.com/joho/godotenv"
	"go-graphql/db"
	"go-graphql/graph"
	"go-graphql/graph/loader"
	"go-graphql/graph/storage"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := getPort()
	mysql, err := db.CreateFromEnv().Open()

	// don't create in here
	//	mysql.MustExec(db.Schema)

	userStorage := storage.NewUserStroage()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	serverWithLoaders := loader.Middleware(userStorage, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", serverWithLoaders)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
