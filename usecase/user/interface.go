package user

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
)

type IUseCase interface {
	GetList(ctx context.Context, req *payload.GetListUserRequest) (*presenter.ListUserResponseWrapper, error)
	GetByID(ctx context.Context, req *payload.GetUserByIDRequest) (*presenter.UserResponseWrapper, error)
	Update(ctx context.Context, req *payload.UpdateUserRequest) (*presenter.UserResponseWrapper, error)
	UpdateInActiveUser(ctx context.Context, req *payload.UpdateUserStatusRequest) (*presenter.UserResponseWrapper, error)
}
