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

func (u userReader) getUsers(ctx context.Context, userIds []int64) ([]*model.User, []error) {

	users, err := u.userStorage.GetUsers(userIds)
	if err != nil {
		return nil, nil
	}

	return users, nil
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID int64) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

// GetUsers returns many users by ids efficiently
func GetUsers(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	loaders := For(ctx)
	fmt.Println("getusers")
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
