package service

import (
	"context"
	"time"

	_ "github.com/lib/pq"
	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/entity"
	"github.com/prapsky/sawitpro/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository  repository.RepositoryInterface
	authService AuthService
}

type UserServiceOptions struct {
	Repository  repository.RepositoryInterface
	AuthService AuthService
}

type Service interface {
	Register(ctx context.Context, input RegisterInput) (uint64, error)
	Login(ctx context.Context, input LoginInput) (LoginOutput, error)
}

func NewUserService(opts UserServiceOptions) *UserService {
	return &UserService{
		repository:  opts.Repository,
		authService: opts.AuthService,
	}
}

type RegisterInput struct {
	PhoneNumber  string
	FullName     string
	PasswordHash string
}

type LoginInput struct {
	PhoneNumber string
	Password    string
}

type LoginOutput struct {
	UserID uint64 `json:"userID"`
	Token  string `json:"token"`
}

func (u *UserService) Register(ctx context.Context, input RegisterInput) (uint64, error) {
	currentTime := time.Now()
	user := entity.User{
		PhoneNumber:  input.PhoneNumber,
		FullName:     input.FullName,
		PasswordHash: input.PasswordHash,
		CreatedAt:    currentTime,
	}

	account, err := u.repository.FindByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		return 0, err
	}

	if account != nil {
		return 0, errors.ErrPhoneNumberAlreadyRegisterd
	}

	userID, err := u.repository.Insert(ctx, user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *UserService) Login(ctx context.Context, input LoginInput) (LoginOutput, error) {
	account, err := u.repository.FindByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		return LoginOutput{}, err
	}

	if account == nil {
		return LoginOutput{}, errors.ErrPhoneNumberNotRegisterd
	}

	inputAttempt := entity.LoginAttempt{
		UserID: account.ID,
	}

	if err := u.comparePasswords(account.PasswordHash, input.Password); err != nil {
		inputAttempt.Success = false
		inputAttempt.AttemptedAt = time.Now()

		if errAttempt := u.repository.InsertLoginAttempts(ctx, inputAttempt); errAttempt != nil {
			return LoginOutput{}, err
		}

		return LoginOutput{}, errors.ErrIncorrectPassword
	}

	inputAttempt.Success = true
	inputAttempt.AttemptedAt = time.Now()
	if errAttempt := u.repository.InsertLoginAttempts(ctx, inputAttempt); errAttempt != nil {
		return LoginOutput{}, err
	}

	account.SuccessfulLogins += 1
	account.LastLoginAt = time.Now()
	if err := u.repository.UpdateSuccessfulLogins(ctx, *account); err != nil {
		return LoginOutput{}, err
	}

	token, err := u.authService.CreateToken(account)
	if err != nil {
		return LoginOutput{}, err
	}

	return LoginOutput{
		UserID: account.ID,
		Token:  token}, nil
}

func (u *UserService) comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
