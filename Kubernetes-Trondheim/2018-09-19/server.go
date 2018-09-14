package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/rs/cors"

	"log"
	"net/http"
	"fmt"
	"math/rand"
)

var locationObject = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Location",
		Description: "An object representing a location for which weather information can be obtained",
		Fields: graphql.Fields{
			"raining": &graphql.Field{
				Description: "An boolean indicating if it's currently raining at the given location",
				Type:        graphql.NewNonNull(graphql.Boolean),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					coordinate := p.Source.(Coordinate)
					fmt.Printf("Checking if it's raning at lat %v, long %v\n", coordinate.Lat, coordinate.Long)
					raining := rand.Intn(2)%2 == 1
					return raining, nil
				},
			},
		},
	},
)

var coordinateInputObject = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "Coordinate",
		Description: "Coordinates consisting of latitude and longitude.",
		Fields: graphql.InputObjectConfigFieldMap{
			"lat": &graphql.InputObjectFieldConfig{
				Type:        graphql.Float,
				Description: "Latitude",
			},
			"long": &graphql.InputObjectFieldConfig{
				Type:        graphql.Float,
				Description: "Longitude",
			},
		}})

type Coordinate struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery", Fields: graphql.Fields{
		"locationFromCoordinates": &graphql.Field{
			Description: "Get location form geo coordinates",
			Args: graphql.FieldConfigArgument{
				"coordinate": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(coordinateInputObject),
					Description: "Coordinate used to identify the location",
				},
			},
			Type: graphql.NewNonNull(locationObject),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				coordinate := p.Args["coordinate"].(map[string]interface{})
				lat := coordinate["lat"].(float64)
				long := coordinate["long"].(float64)
				return Coordinate{Lat: lat, Long: long}, nil
			},
		},
	},
	Description: `Root query of the Mito.ai weather service.`})

var schemaConfig = graphql.SchemaConfig{
	Query: rootQuery,
}

func main() {
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	mux := http.NewServeMux()
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})
	mux.Handle("/graphql", h)
	println("Listening on port 1337")
	http.ListenAndServe(":1337", cors.Default().Handler(mux))
}
