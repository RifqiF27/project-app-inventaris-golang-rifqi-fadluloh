package service

import (
	"inventaris/model"
	"inventaris/repository"

	"github.com/google/uuid"
)

type UserService struct {
	RepoUser    repository.UserRepository
	RepoSession repository.SessionRepository
}

func NewUserService(repoUser repository.UserRepository, repoSession repository.SessionRepository) UserService {
	return UserService{repoUser, repoSession}
}

func (us *UserService) LoginService(user model.User) (*model.User, error) {

	users, _ := us.RepoUser.GetUserLogin(user)

	return users, nil
}

func (s *UserService) RegisterUser(user model.User) error {
	return s.RepoUser.Create(&user)
}

func (s *UserService) CreateSession(userID int) (uuid.UUID, error) {
	return s.RepoSession.CreateSession(userID)
}

func (s *UserService) GetUserIDBySessionID(sessionID string) (int, error) {
	return s.RepoSession.GetUserIDBySessionID(sessionID)
}

func (s *UserService) DeleteSession(sessionID string) error {
	return s.RepoSession.DeleteSession(sessionID)
}
