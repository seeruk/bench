package graphql

import (
	"context"
	"testing"

	graphqlb "github.com/graph-gophers/graphql-go"
	graphqlbsw "github.com/graph-gophers/graphql-go/example/starwars"
	graphqla "github.com/graphql-go/graphql"
	graphqlasw "github.com/graphql-go/relay/examples/starwars"
)

func BenchmarkGraphQLGoGraphQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		query := `{
			rebels {
				ships {
					edges {
						node {
							id
							name
						}
					}
				}
			}
		}`

		params := graphqla.Params{
			Schema:        graphqlasw.Schema,
			RequestString: query,
		}

		r := graphqla.Do(params)

		if r.HasErrors() {
			for _, err := range r.Errors {
				b.Log(err)
			}

			b.Error("error(s) occurred")
		}
	}
}

func BenchmarkGraphQLGophersGraphQLGo(b *testing.B) {
	schema := graphqlb.MustParseSchema(graphqlbsw.Schema, &graphqlbsw.Resolver{})

	query := `{
		hero {
			friends {
				name
				appearsIn
			}
		}
	}`

	var res *graphqlb.Response

	for i := 0; i < b.N; i++ {
		res = schema.Exec(context.Background(), query, "", nil)

		if len(res.Errors) > 0 {
			for _, err := range res.Errors {
				b.Log(err)
			}

			b.Error("error(s) occurred")
		}
	}
}
