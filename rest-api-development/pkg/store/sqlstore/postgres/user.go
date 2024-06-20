package postgres

import "github.com/alikazai/golearn/pkg/domain"

func (q *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1
	return user, nil
}
