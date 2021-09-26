package application

import (
	"context"

	"github.com/kawamou/my-layered-architecture/domain"
	"github.com/pkg/errors"
)

type FindUserService struct {
	r domain.UserRepository
}

type FindUserServiceInput struct {
	ID string
}

type FindUserServiceOutput struct {
	User domain.User
}

func NewFindUserService(repo domain.UserRepository) *FindUserService {
	return &FindUserService{
		r: repo,
	}
}

func (f *FindUserService) Handle(ctx context.Context, input FindUserServiceInput) (FindUserServiceOutput, error) {
	user, err := f.r.FindById(ctx, input.ID)
	if err != nil {
		return FindUserServiceOutput{}, errors.Wrap(err, "")
	}
	return FindUserServiceOutput{User: *user}, nil
}
