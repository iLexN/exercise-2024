package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"crypto/rand"
	"fmt"
	"go-graphql/graph/loader"
	"go-graphql/graph/model"
	"math/big"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))

	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", randNumber),
		Text:   input.Text,
		UserID: input.UserID,
		//		Done: false,
	}

	// dependency, memory as db
	r.todos = append(r.todos, todo)

	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user, err := loader.GetUser(ctx, obj.UserID)

	if err != nil {
		return nil, err
	}
	return user, nil
	//	return &model.User{
	//		Name: "gi" + obj.ID,
	//		ID:   obj.ID,
	//	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
