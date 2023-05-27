package storage

import (
	"auth/infra/headers"
	"auth/infra/realisations/user"
	"auth/logger"
	"errors"
	"sync"
)

type MemoryStorage struct {
	users map[string]*user.User
	mu    sync.RWMutex
}

func New(l logger.Logger) (*MemoryStorage, error) {
	return &MemoryStorage{
		users: make(map[string]*user.User),
	}, nil
}

func (s *MemoryStorage) FindUser(login string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.users[login]
	return ok, nil
}

func (s *MemoryStorage) GetUser(login string, password string) (headers.User, error) {
	s.mu.RLock()
	user, ok := s.users[login]
	s.mu.RUnlock()

	if !ok {
		return nil, errors.New("GetUser: user not found")
	}
	if password != user.GetPassword() {
		return nil, errors.New("GetUser: invalid password")
	}
	return user, nil
}

func (s *MemoryStorage) InsertUser(newUser headers.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[newUser.GetLogin()]; ok {
		return errors.New("InsertUser: user already exists")
	}
	s.users[newUser.GetLogin()] = user.NewUser(newUser.GetLogin(), newUser.GetPassword())
	return nil
}
