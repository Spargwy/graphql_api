package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gql_app/graph/generated"
	"gql_app/graph/model"
	"gql_app/storage"
	"log"
)

var usersCodes = make(map[string]string)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {
	code := generateRandomCode(4)
	log.Print(code)
	usersCodes[input.Phone] = code
	payload := &model.ErrorPayload{Message: "null"}
	return payload, nil
}

func (r *mutationResolver) SignInByCode(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload, error) {
	if input.Code == usersCodes[input.Phone] {
		user, err := storage.SelectUserByPhone(input.Phone)
		if err != nil {
			log.Print("SelectUserByPhone error: ", err)
			return nil, err
		}
		return model.SignInPayload{Viewer: &model.Viewer{User: &user}, Token: "123"}, nil
	} else {
		return model.ErrorPayload{Message: "Invalid code"}, nil
	}
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var err error
	r.products, err = storage.SelectProducts()
	if err != nil {
		return r.products, err
	}
	return r.products, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {
	log.Print("Viewer not implemented")
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
