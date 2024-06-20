package postgres

import (
	"testing"

	"github.com/alikazai/golearn/pkg/domain"
)

func TestCreateUser(t *testing.T) {

	pStore := NewPostgresStore(nil)

	user := &domain.User{
		Email:    "test@test.com",
		Password: "password",
		Name:     "John Doe",
	}

	createdUser, err := pStore.CreateUser(user)

	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Errorf("Want id not to be zero")
	}

}
