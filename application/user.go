package application

import (
	"context"

	"github.com/takashabe/go-ddd-sample/domain"
	"github.com/takashabe/go-ddd-sample/domain/repository"
)

// UserInteractor provides use-case 
type UserInteractor struct {
	repository repository.UserRepository
}

// GetUser returns user
func (i UserInteractor) GetUser(ctx context.Context, id int) (*domain.User, error) {
	return i.repository.Get(ctx, id)
}

// GetUsers returns user list
func (i UserInteractor) GetUsers(ctx context.Context) ([]*domain.User, error) {
	return i.repository.GetAll(ctx)
}

// AddUser saves new user
func (i UserInteractor) AddUser(ctx context.Context, name string) error {
	u, err := domain.NewUser(name)
	if err != nil {
		return err
	}
	return i.repository.Save(ctx, u)
}
