package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql_app/auth"
	"gql_app/graph/generated"
	"gql_app/graph/model"
	"log"
)

// Просто, чтобы не усложнять. По-хорошему, конечно, Redis юзать
var usersCodes = make(map[string]string)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {
	err := r.SendCode(input.Phone)

	if err != nil {
		return &model.ErrorPayload{Message: "Cant send message"}, err
	}

	payload := &model.ErrorPayload{Message: "null"}

	return payload, nil
}

func (r *mutationResolver) SignInByCode(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload, error) {
	if input.Code != usersCodes[input.Phone] {
		return model.ErrorPayload{Message: "Invalid code"}, nil
	}

	user, err := r.SelectUserByPhone(input.Phone)

	if err != nil {
		log.Print("SelectUserByPhone error: ", err)
		return nil, err
	}

	token, err := generateJWT(user.ID)

	if err != nil {
		return model.ErrorPayload{Message: "Internal error"}, err
	}

	delete(usersCodes, input.Phone)

	return model.SignInPayload{Viewer: &model.Viewer{User: &user}, Token: token}, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var err error

	r.products, err = r.SelectProducts()
	if err != nil {
		return r.products, err
	}

	return r.products, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {
	viewer := auth.ForContext(ctx)
	if viewer == nil {
		return &model.Viewer{}, fmt.Errorf("viewer is not found")
	}

	return viewer, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
