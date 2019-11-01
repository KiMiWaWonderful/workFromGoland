package services

import (
	"errors"
	"github.com/kataras/iris/_examples/mvc/login/datamodels"
	"github.com/kataras/iris/_examples/mvc/login/repositories"
)

type UserService interface {
	GetAll() []datamodels.User
	GetByID(id int64)(datamodels.User, bool)
	GetByUsernameAndPassword(username, userPassword string) (datamodels.User,bool)
	DeleteByID(id int64) bool

	Update(id int64, user datamodels.User) (datamodels.User,error)
	UpdatePassword(id int64, newPassword string) (datamodels.User,error)
	UpdateUsername(id int64, newUsername string) (datamodels.User,error)

	Create(userPassword string, user datamodels.User) (datamodels.User,error)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo:repo,
	}
}

type userService struct {
	repo repositories.UserRepository
}

func (s *userService) GetAll() []datamodels.User {
	return s.repo.SelectMany(func(_ datamodels.User) bool {
		return true
	},-1)
}

func (s *userService) GetByID(id int64) (datamodels.User, bool) {
	return s.repo.Select(func(m datamodels.User) bool {
		return m.ID == id
	})
}

func (s *userService) GetByUsernameAndPassword(username, userPassword string) (datamodels.User, bool) {
	if username == "" || userPassword == ""{
		return datamodels.User{},false
	}

	return s.repo.Select(func(m datamodels.User) bool {
		if m.Username == username {
			hashed := m.HashedPassword
			if ok,_ := datamodels.ValidatePassword(userPassword,hashed); ok{
				return true
			}
		}
		return false
	})
}

func (s *userService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.User) bool {
		return m.ID == id
	}, 1)
}

func (s *userService) Update(id int64, user datamodels.User) (datamodels.User, error) {
	user.ID = id
	return s.repo.InsertOrUpdate(user)
}

func (s *userService) UpdatePassword(id int64, newPassword string) (datamodels.User, error) {
	hashed, err := datamodels.GeneratePassword(newPassword)
	if err != nil{
		return datamodels.User{},err
	}

	return s.Update(id, datamodels.User{
		HashedPassword:hashed,
	})
}

func (s *userService) UpdateUsername(id int64, newUsername string) (datamodels.User, error) {
	return s.Update(id, datamodels.User{
		Username: newUsername,
	})
}

func (s *userService) Create(userPassword string, user datamodels.User) (datamodels.User, error) {
	if user.ID > 0 || userPassword == "" || user.Firstname == "" || user.Username == "" {
		return datamodels.User{}, errors.New("unable to create this user")
	}

	hashed, err := datamodels.GeneratePassword(userPassword)
	if err != nil {
		return datamodels.User{}, err
	}
	user.HashedPassword = hashed

	return s.repo.InsertOrUpdate(user)
}
