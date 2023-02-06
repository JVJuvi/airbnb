package payload

import (
	validation "go01-airbnb/client/validator"
	"go01-airbnb/codetype"
	myerror "go01-airbnb/util/error"
	"strings"
)

type CreateLocationRequest struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}

type GetLocationByIDRequest struct {
	ID int64 `json:"-"`
}

type GetListLocationRequest struct {
	codetype.Paginator
}

func (g *GetListLocationRequest) Format() {
	g.Paginator.Format()
}

type UpdateLocationRequest struct {
	ID      int64   `json:"-"`
	Country *string `json:"country"`
	State   *string `json:"state"`
	City    *string `json:"city"`
}

type DeleteLocationRequest struct {
	ID int64 `json:"-"`
}

func (u *CreateLocationRequest) Validate() error {
	u.Country = strings.TrimSpace(u.Country)
	u.State = strings.TrimSpace(u.State)
	u.City = strings.TrimSpace(u.City)
	err := validation.GetValidator().Struct(u)
	if err != nil {
		return myerror.ErrStructValidation(err)
	}

	return nil
}
