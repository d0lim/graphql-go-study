package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/d0lim/graphql-go-study/schema"
)

func init() {
	todo1 := schema.Todo{ID: "a", Text: "A todo not to forget", Done: false}
	todo2 := schema.Todo{ID: "b", Text: "This is the most important", Done: false}
	todo3 := schema.Todo{ID: "c", Text: "Please do this or else", Done: false}
	schema.TodoList = append(schema.TodoList, todo1, todo2, todo3)

	rand.Seed(time.Now().UnixNano())
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {

	h := handler.New(&handler.Config{
		Schema:     &schema.TodoSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/graphql", h)
	http.Handle("/", fs)

	http.ListenAndServe(":5000", nil)
}
