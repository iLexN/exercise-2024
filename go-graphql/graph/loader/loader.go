package loader

import (
	"context"
	"github.com/vikstrous/dataloadgen"
	"go-graphql/graph/model"
	"net/http"
	"time"
)

// https://github.com/vikstrous/dataloadgen-example/blob/master/graph/loader/loader.go

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type userReader struct {
	userStorage *string
}

// Loaders wrap your data loaders to inject via middleware
func (u userReader) getUsers(ctx context.Context, userIds []string) ([]*model.User, []error) {

	users := make([]*model.User, 0, len(userIds))

	for _, userId := range userIds {
		user := &model.User{
			Name: "mu" + userId,
			ID:   userId,
		}
		users = append(users, user)
	}

	return users, nil
}

type Loaders struct {
	UserLoader *dataloadgen.Loader[string, *model.User]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(s *string) *Loaders {
	// define the data loader
	userReader := &userReader{userStorage: s}
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(userReader.getUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}

func Middleware(userStorage *string, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Note that the loaders are being created per-request. This is important because they contain caching and batching logic that must be request-scoped.
		loaders := NewLoaders(userStorage)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loaders))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser returns single user by id efficiently
func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.Load(ctx, userID)
}

// GetUsers returns many users by ids efficiently
func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, error) {
	loaders := For(ctx)
	return loaders.UserLoader.LoadAll(ctx, userIDs)
}
