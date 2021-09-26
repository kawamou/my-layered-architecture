package application

import (
	"context"

	"github.com/kawamou/my-layered-architecture/domain"
	"github.com/pkg/errors"
)

// CreateUserService ...
type CreateUserService struct {
	r domain.UserRepository
}

type CreateUserServiceInput struct {
	ID   string
	Name string
	Age  int
}

func NewCreateUserService(repo domain.UserRepository) *CreateUserService {
	return &CreateUserService{
		r: repo,
	}
}

func (c *CreateUserService) Handle(ctx context.Context, input CreateUserServiceInput) error {
	user := domain.NewUser(input.ID, input.Name, input.Age)
	if err := c.r.Add(ctx, *user); err != nil {
		errors.Wrap(err, "")
	}
	return nil
}
