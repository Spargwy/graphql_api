package resolvers

import (
	"gql_app/graph/model"
	"gql_app/graph/resolvers/storage"

	"github.com/twilio/twilio-go"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	products []*model.Product
	storage.Psql
	*twilio.RestClient
}
