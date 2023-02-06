package place

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

func (r *pgRepository) Create(ctx context.Context, data *model.Place) error {
	if err := r.getDB(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) GetByID(ctx context.Context, id int64) (*model.Place, error) {
	var place model.Place

	err := r.getDB(ctx).Where("id = ?", id).First(&place).Error
	if err != nil {
		return nil, err
	}

	return &place, nil
}

func (r *pgRepository) GetList(
	ctx context.Context,
	paginator *codetype.Paginator,
) ([]model.Place, int64, error) {
	var (
		db     = r.getDB(ctx).Model(&model.Place{})
		data   = make([]model.Place, 0)
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

func (r *pgRepository) Delete(ctx context.Context, data *model.Place) error {
	err := r.getDB(ctx).Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *pgRepository) Update(ctx context.Context, data *model.Place) error {
	return r.getDB(ctx).Save(data).Error
}
