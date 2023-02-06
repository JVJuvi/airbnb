package auth

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
)

type IUseCase interface {
	Register(ctx context.Context, req *payload.RegisterRequest) error
	Login(ctx context.Context, req *payload.LoginRequest) (*presenter.LoginResponseWrapper, error)
}
