package main

import (
	"fmt"
	"graphgo/gql"
	"graphgo/postgres"
	"graphgo/server"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(postgres.ConnString("graphgo_database", 5432, "docker", "docker", "go_graph_db"))
	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery.Query})
	if err != nil {
		fmt.Println("Error creating scehma: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	router.Post("/graphql", s.GraphQL())

	return router, db
}
