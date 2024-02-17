package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"go-graphql/graph/model"
	"go-graphql/graph/storage"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos       []*model.Todo
	UserStorage *storage.UserStorage
}

func (r *Resolver) GetTodosByUserId(user *model.User) ([]*model.Todo, error) {
	var userTodos []*model.Todo

	// Iterate over all todos to find the ones associated with the given user ID
	for _, todo := range r.todos {
		if todo.UserID == user.ID {
			userTodos = append(userTodos, todo)
		}
	}

	return userTodos, nil
}
