package location

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
)

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListLocationRequest,
) (*presenter.ListLocationResponseWrapper, error) {
	req.Format()

	locations, total, err := u.LocationRepo.GetList(ctx, &req.Paginator)
	if err != nil {
		return nil, myerror.ErrLocationGet(err)
	}

	return &presenter.ListLocationResponseWrapper{
		Location: locations,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}, nil
}
