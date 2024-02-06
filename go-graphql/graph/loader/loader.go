package loader

import (
	"context"
	"net/http"
	"time"

	"github.com/vikstrous/dataloadgen"
	"go-graphql/graph/model"
	"go-graphql/graph/storage"
)

// https://github.com/vikstrous/dataloadgen-example/blob/master/graph/loader/loader.go

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	UserLoader *dataloadgen.Loader[int64, *model.User]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(s *storage.UserStorage) *Loaders {
	// define the data loader
	userReader := &userReader{userStorage: s}
	return &Loaders{
		UserLoader: dataloadgen.NewLoader(userReader.getUsers, dataloadgen.WithWait(time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(userStorage *storage.UserStorage, next http.Handler) http.Handler {
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
