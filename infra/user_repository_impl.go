package infra

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/kawamou/my-layered-architecture/domain"
	"github.com/pkg/errors"
)

const collectionName = "users"

type UserRepositoryImpl struct {
	client *firestore.Client
}

func NewUserRepositoryImpl(ctx context.Context, client *firestore.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		client: client,
	}
}

func (u *UserRepositoryImpl) Add(ctx context.Context, user domain.User) error {
	_, err := u.client.Collection(collectionName).Doc(user.ID).Set(ctx, user)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User
	snap, err := u.client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if err := snap.DataTo(&user); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &user, nil
}
