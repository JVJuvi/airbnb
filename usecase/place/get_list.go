package place

import (
	"context"
	"go01-airbnb/payload"
	"go01-airbnb/presenter"
	myerror "go01-airbnb/util/error"
)

func (u *UseCase) GetList(
	ctx context.Context,
	req *payload.GetListPlaceRequest,
) (*presenter.ListPlaceResponseWrapper, error) {
	req.Format()

	places, total, err := u.PlaceRepo.GetList(ctx, &req.Paginator)
	if err != nil {
		return nil, myerror.ErrLocationGet(err)
	}

	return &presenter.ListPlaceResponseWrapper{
		Place: places,
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}, nil
}
