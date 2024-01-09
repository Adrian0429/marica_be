package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	IframesRepository interface {
		CreateIframes(ctx context.Context, iframe entities.Iframes) (entities.Iframes, error)
	}

	iframeRepository struct {
		db *gorm.DB
	}
)

func NewIframesRepository(db *gorm.DB) IframesRepository {
	return &iframeRepository{
		db: db,
	}
}

func (r *iframeRepository) CreateIframes(ctx context.Context, iframe entities.Iframes) (entities.Iframes, error) {
	if err := r.db.Create(&iframe).Error; err != nil {
		return entities.Iframes{}, err
	}

	return iframe, nil
}

func (r *iframeRepository) DeleteIframes(ctx context.Context, iframe entities.Iframes) (entities.Iframes, error) {
	return iframe, nil
}
