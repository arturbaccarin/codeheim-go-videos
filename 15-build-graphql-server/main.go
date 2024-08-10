package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// 1 Define Structs: these structures represent the data
type Blog struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal(err)
	}
}

// 2 Create GraphQL types: define graphql object types that correspond to the data structures
func createBlogType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Blog",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

// 3 Define GraphQL Schema: define the queries that can be run on the server
func queryType(blogType *graphql.Object) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"blogs": &graphql.Field{
					Type: graphql.NewList(blogType),
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Read limit
						limit, _ := p.Args["limit"].(int)
						if limit <= 0 || limit > 20 {
							limit = 10
						}

						// Read offset
						offset, _ := p.Args["offset"].(int)
						if offset <= 0 || offset > 20 {
							offset = 10
						}

						return getBlogs(limit, offset)
					},
				},
			},
		},
	)
}

func getBlogs(limit, offset int) ([]Blog, error) {
	var blogs []Blog
	rows, err := db.Query("SELECT id, title, content FROM blog limit" + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Blog
		err := rows.Scan(&b.ID, &b.Title, &b.Content)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, b)
	}

	return blogs, nil
}

func main() {
	initDB()

	blogType := createBlogType()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType(blogType),
		},
	)

	if err != nil {
		log.Fatal("failed to create new schema, error: %v", err)
	}

	handler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// 4 Setup HTTP Server: register the GraphQL handler with an HTTP route.
	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
