package storage

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go-graphql/graph/model"
	"log"
	"sync"
)

var (
	notFoundError = errors.New("not found")
	//	existsError   = errors.New("exists")
)

type UserStorage struct {
	lock sync.Mutex
	data map[string]model.User
	db   *sqlx.DB
}

func NewUserStroage(db *sqlx.DB) *UserStorage {
	return &UserStorage{
		data: make(map[string]model.User),
		db:   db,
	}
}

func (u *UserStorage) Get(userId int64) (*model.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	var user model.User

	log.Printf("Debug: UserStorage Get: %d\n", userId)

	err := u.db.Get(&user, "SELECT * FROM users WHERE id = ?", userId)
	if err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func (u *UserStorage) Put(user model.NewUser) (*model.User, error) {
	u.lock.Lock()
	defer u.lock.Unlock()

	log.Printf("Debug: UserStorage Put - Name: %d\n", user.Name)

	res, err := u.db.NamedExec("INSERT INTO `users` (`name`) VALUES (:name)", user)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		ID:   id,
		Name: user.Name,
	}

	return newUser, nil
}
