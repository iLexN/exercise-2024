package loader

import (
	"context"
	"fmt"

	"go-graphql/graph/model"
	"go-graphql/graph/storage"
)

type userReader struct {
	userStorage *storage.UserStorage
}

func (u userReader) getUsers(ctx context.Context, userIds []string) ([]*model.User, []error) {

	users := make([]*model.User, 0, len(userIds))

	fmt.Printf("UserStorage getUsers %d\n", len(userIds))

	for _, userId := range userIds {
		user := &model.User{
			Name: "mu" + userId,
			ID:   userId,
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	fmt.Println("getuser")
	return loaders.UserLoader.Load(ctx, userID)
}

// GetUsers returns many users by ids efficiently
func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	fmt.Println("getusers")
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
