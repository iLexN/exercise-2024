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

	//todo: https://cloud.tencent.com/developer/article/1997275

	users := make([]*model.User, 0, len(userIds))

	fmt.Printf("UserStorage getUsers %d\n", len(userIds))

	for _, userId := range userIds {

		user := &model.User{
			Name: "mu",
			ID:   userId,
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID int64) (*model.User, error) {
	loaders := For(ctx)
	fmt.Println("getuser")
	return loaders.UserLoader.Load(ctx, userID)
}

// GetUsers returns many users by ids efficiently
func GetUsers(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	loaders := For(ctx)
	fmt.Println("getusers")
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
