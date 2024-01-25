package service

import (
	"context"
	"time"

	_ "github.com/lib/pq"
	"github.com/prapsky/sawitpro/common/errors"
	"github.com/prapsky/sawitpro/entity"
	"github.com/prapsky/sawitpro/generated"
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
	Register(ctx context.Context, input RegisterInput) (generated.RegisterResponse, error)
	Login(ctx context.Context, input LoginInput) (generated.LoginResponse, error)
	GetProfile(ctx context.Context, token string) (generated.ProfileResponse, error)
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

type GetProfileOutput struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

func (u *UserService) Register(ctx context.Context, input RegisterInput) (generated.RegisterResponse, error) {
	currentTime := time.Now()
	user := entity.User{
		PhoneNumber:  input.PhoneNumber,
		FullName:     input.FullName,
		PasswordHash: input.PasswordHash,
		CreatedAt:    currentTime,
	}

	account, err := u.repository.FindByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		return generated.RegisterResponse{}, err
	}

	if account != nil {
		return generated.RegisterResponse{}, errors.ErrPhoneNumberAlreadyRegisterd
	}

	userID, err := u.repository.Insert(ctx, user)
	if err != nil {
		return generated.RegisterResponse{}, err
	}

	data := &generated.RegisterResponseData{
		UserID: &userID,
	}

	return generated.RegisterResponse{
		Data: data}, nil
}

func (u *UserService) Login(ctx context.Context, input LoginInput) (generated.LoginResponse, error) {
	account, err := u.repository.FindByPhoneNumber(ctx, input.PhoneNumber)
	if err != nil {
		return generated.LoginResponse{}, err
	}

	if account == nil {
		return generated.LoginResponse{}, errors.ErrPhoneNumberNotRegisterd
	}

	inputAttempt := entity.LoginAttempt{
		UserID: account.ID,
	}

	if err := u.comparePasswords(account.PasswordHash, input.Password); err != nil {
		inputAttempt.Success = false
		inputAttempt.AttemptedAt = time.Now()

		if errAttempt := u.repository.InsertLoginAttempts(ctx, inputAttempt); errAttempt != nil {
			return generated.LoginResponse{}, err
		}

		return generated.LoginResponse{}, errors.ErrIncorrectPassword
	}

	inputAttempt.Success = true
	inputAttempt.AttemptedAt = time.Now()
	if errAttempt := u.repository.InsertLoginAttempts(ctx, inputAttempt); errAttempt != nil {
		return generated.LoginResponse{}, err
	}

	account.SuccessfulLogins += 1
	account.LastLoginAt = time.Now()
	if err := u.repository.UpdateSuccessfulLogins(ctx, *account); err != nil {
		return generated.LoginResponse{}, err
	}

	token, err := u.authService.CreateToken(account)
	if err != nil {
		return generated.LoginResponse{}, err
	}

	data := &generated.LoginResponseData{
		Token:  &token,
		UserID: &account.ID,
	}

	return generated.LoginResponse{
		Data: data}, nil
}

func (u *UserService) comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserService) GetProfile(ctx context.Context, token string) (generated.ProfileResponse, error) {
	userID, err := u.authService.ValidateToken(token)
	if err != nil {
		return generated.ProfileResponse{}, err
	}

	account, err := u.repository.FindByID(ctx, userID)
	if err != nil {
		return generated.ProfileResponse{}, err
	}

	if account == nil {
		return generated.ProfileResponse{}, errors.ErrPhoneNumberNotRegisterd
	}

	data := &generated.ProfileResponseData{
		FullName:    &account.FullName,
		PhoneNumber: &account.PhoneNumber,
	}

	return generated.ProfileResponse{
		Data: data}, nil
}
