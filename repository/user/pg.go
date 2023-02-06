package user

import (
	"context"
	"go01-airbnb/codetype"
	"go01-airbnb/model"
	"gorm.io/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (r pgRepository) Create(ctx context.Context, data *model.User) error {
	return r.getDB(ctx).Create(data).Error
}

func (r *pgRepository) GetByInterface(ctx context.Context, conditions interface{}) (*model.User, error) {
	var user model.User
	if err := r.getDB(ctx).Where(conditions).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *pgRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User

	err := r.getDB(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *pgRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.getDB(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *pgRepository) GetList(
	ctx context.Context,
	paginator *codetype.Paginator,
) ([]model.User, int64, error) {
	var (
		db     = r.getDB(ctx).Model(&model.User{})
		data   = make([]model.User, 0)
		offset int
		total  int64
	)

	if paginator.Page != 1 {
		offset = paginator.Limit * (paginator.Page - 1)
	}

	if paginator.Limit != -1 {
		if err := db.Count(&total).Error; err != nil {
			return nil, 0, err
		}
	}

	err := db.Limit(paginator.Limit).Offset(offset).Find(&data).Error
	if err != nil {
		return nil, 0, err
	}

	if paginator.Limit == -1 {
		total = int64(len(data))
	}

	return data, total, nil
}

func (r *pgRepository) Update(ctx context.Context, data *model.User) error {
	return r.getDB(ctx).Save(data).Error
}
