// This file contains the service implementation layer.
package service

import (
	"context"
	"time"

	_ "github.com/lib/pq"
	"github.com/prapsky/sawitpro/entity"
	"github.com/prapsky/sawitpro/repository"
)

type UserService struct {
	repository repository.RepositoryInterface
}

type UserServiceOptions struct {
	Repository repository.RepositoryInterface
}

type Service interface {
	Register(ctx context.Context, input RegisterInput) (uint64, error)
}

func NewUserService(opts UserServiceOptions) *UserService {
	return &UserService{
		repository: opts.Repository,
	}
}

type RegisterInput struct {
	PhoneNumber  string
	FullName     string
	PasswordHash string
}

func (u *UserService) Register(ctx context.Context, input RegisterInput) (uint64, error) {

	currentTime := time.Now()
	user := entity.User{
		PhoneNumber:  input.PhoneNumber,
		FullName:     input.FullName,
		PasswordHash: input.PasswordHash,
		CreatedAt:    currentTime,
	}

	userID, err := u.repository.Insert(ctx, user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
