package service

import (
    "context"
    "errors"
    "user-service/internal/adapter/repository"
    "user-service/internal/core/domain/entity"
    "user-service/utils/conv"

    "github.com/rs/zerolog/log"
)

type UserServiceInterface interface {
    SignIn(ctx context.Context, email string, password string) (*entity.UserEntity, error)
}

type userService struct {
    repo repository.UserRepositoryInterface
}

func (u *userService) SignIn(ctx context.Context, email string, password string) (*entity.UserEntity, error) {
    user, err := u.repo.GetUserByEmail(ctx, email)
    if err != nil {
        log.Error().Err(err).Msg("[UserService-1] SignIn: failed to get user")
        return nil, err
    }
    
    if checkPass := conv.CheckPasswordHash(password, user.Password); !checkPass {
        err = errors.New("password is incorrect")
        log.Error().Err(err).Msg("[UserService-2] SignIn: invalid password")
        return nil, err
    }
    
    return user, nil
}

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
    return &userService{repo: repo}
}