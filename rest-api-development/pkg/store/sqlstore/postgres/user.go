package postgres

import (
	"github.com/alikazai/golearn/pkg/common"
	"github.com/alikazai/golearn/pkg/domain"
)

func (q *postgresStore) CreateUser(user *domain.User) (*domain.User, error) {
	user.ID = 1
	user.Password, _ = common.PasswordHash(user.Password)
	return user, nil
}
