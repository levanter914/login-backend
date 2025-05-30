package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.72

import (
	"context"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/levanter914/login-page/backend/graph/model"
	"os"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthPayload, error) {
	if email != "user@example.com" || password != "password123" {
		return nil, errors.New("invalid credentials")
	}

	name := "John Doe"
	user := &model.User{
		ID:    "1",
		Email: "user@example.com",
		Name:  &name,
	}


	claims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, errors.New("could not sign JWT")
	}

	return &model.AuthPayload{
		Token: signedToken,
		User:  user,
	}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userID, ok := GetUserIDFromContext(ctx)
	if !ok {
		return nil, errors.New("unauthenticated")
	}

	if userID != "1" {
		return nil, errors.New("user not found")
	}

	name := "John Doe"
	return &model.User{
		ID:    "1",
		Email: "user@example.com",
		Name:  &name,
	}, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
