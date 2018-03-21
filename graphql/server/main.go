package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/example/starwars"
)

type Request struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello, World!")
	})

	r.Get("/graphiql", func(w http.ResponseWriter, r *http.Request) {
		w.Write(graphiql)
	})

	r.Post("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var request Request

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		schema := graphql.MustParseSchema(starwars.Schema, &starwars.Resolver{})
		result := schema.Exec(r.Context(), request.Query, "", nil)

		w.Write(result.Data)
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}

var graphiql = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.css"/>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function fetchGQL(params) {
				return fetch("/graphql", {
					method: "post",
					body: JSON.stringify(params),
					credentials: "include",
				}).then(function (resp) {
					return resp.text();
				}).then(function (body) {
					try {
						return JSON.parse(body);
					} catch (error) {
						return body;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: fetchGQL}),
				document.getElementById("graphiql")
			)
		</script>
	</body>
</html>
`)
