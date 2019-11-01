package repositories

import (
	"errors"
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"sync"
)

type Query func(user datamodels.User) bool

type UserRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)

	Select(query Query) (user datamodels.User, found bool)
	SelectMany(query Query, limit int) (results []datamodels.User)

	InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error)
	Delete(query Query, limit int) (deleted bool)
}

func NewUserRepository(source map[int64]datamodels.User) UserRepository {
	return &userMemoryRepository{source:source}
}

type userMemoryRepository struct {
	source map[int64]datamodels.User
	mu sync.RWMutex
}

const   (
	ReadOnlyMode = iota
	ReadWriteMode
)

func (r *userMemoryRepository) Exec(query Query, action Query, limit int, mode int) (ok bool) {
	loops := 0

	if mode == ReadOnlyMode{
		r.mu.RLock()
		defer r.mu.RUnlock()
	}else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, user := range r.source{
		ok = query(user)
		if ok {
			if action(user){
				loops ++
				if limit >= loops{
					break
				}
			}
		}
	}
	return
}

func (r *userMemoryRepository) Select(query Query) (user datamodels.User, found bool) {
	found = r.Exec(query, func(m datamodels.User) bool {
		user = m
		return true
	},1,ReadOnlyMode)

	if !found{
		user = datamodels.User{}
	}
	return
}

func (r *userMemoryRepository) SelectMany(query Query, limit int) (results []datamodels.User) {
	r.Exec(query, func(m datamodels.User) bool {
		results = append(results,m)
		return true
	},limit,ReadOnlyMode)
	return
}

func (r *userMemoryRepository) InsertOrUpdate(user datamodels.User) (updatedUser datamodels.User, err error) {
	id := user.ID

	if id == 0{
		var lastID int64
		r.mu.RLock()
		for _, item := range r.source{
			if item.ID > lastID{
				lastID = item.ID
			}
		}
		r.mu.RUnlock()

		id = lastID + 1
		user.ID = id

		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()

		return user,nil
	}

	current,exists := r.Select(func(m datamodels.User) bool {
		return m.ID == id
	})

	if !exists{
		return datamodels.User{}, errors.New("failed to update a nonexistent user")
	}

	if user.Username != "" {
		current.Username = user.Username
	}

	if user.Firstname != "" {
		current.Firstname = user.Firstname
	}

	// map-specific thing
	r.mu.Lock()
	r.source[id] = current
	r.mu.Unlock()

	return user, nil
}

func (r *userMemoryRepository) Delete(query Query, limit int) (deleted bool) {
	return r.Exec(query, func(m datamodels.User) bool {
		delete(r.source,m.ID)
		return true
	},limit,ReadWriteMode)
}
