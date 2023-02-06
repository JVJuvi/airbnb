package user

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
)

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListUserRequest,
) (*presenter.ListUserResponseWrapper, error) {
	req.Format()

	users, total, err := u.UserRepo.GetList(ctx, &req.Paginator)
	if err != nil {
		return nil, myerror.ErrUserGet(err)
	}

	return &presenter.ListUserResponseWrapper{
		User: users,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}, nil
}
