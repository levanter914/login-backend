package graph

import (
	"context"
	"errors"
	"github.com/levanter914/login-backend/graph"
	"github.com/levanter914/login-backend.git/graph/model"
)

// Mutation resolver
type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthPayload, error) {
	msg := "Login successful"
	if email == "user@example.com" && password == "password123" {
		return &model.AuthPayload{
			Token:   "sample-jwt-token",
			Message: &msg,
		}, nil
	}
	return nil, errors.New("invalid credentials")
}

// Query resolver
type queryResolver struct{ *Resolver }

func (r *queryResolver) Status(ctx context.Context) (string, error) {
	return "Server is running", nil
}

// Resolver struct holds both Query and Mutation resolvers
type Resolver struct{}

func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
