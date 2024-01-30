package storage

import (
	"errors"
	"fmt"
	"go-graphql/graph/model"
	"sync"
	"time"
)

var (
	notFoundError = errors.New("not found")
	//	existsError   = errors.New("exists")
)

type UserStorage struct {
	lock sync.Mutex
	data map[string]model.User
}

func NewUserStroage() *UserStorage {
	return &UserStorage{
		data: make(map[string]model.User),
	}
}

func (u *UserStorage) Get(userId string) (model.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	fmt.Println("user get")
	time.Sleep(time.Second)

	user, ok := u.data[userId]
	if !ok {
		return model.User{}, notFoundError
	}

	return user, nil
}
