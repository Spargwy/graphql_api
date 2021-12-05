package resolvers

import (
	"gql_app/graph/model"

	"github.com/go-pg/pg/v10"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	products []*model.Product
	DB       *pg.DB
}
