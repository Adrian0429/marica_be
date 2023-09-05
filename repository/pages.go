package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	PagesRepository interface {
		CreatePages(ctx context.Context, proofOfDamage entities.Pages) (entities.Pages, error)
	}

	pagesRepository struct {
		db *gorm.DB
	}
)

func NewPagesRepository(db *gorm.DB) *pagesRepository {
	return &pagesRepository{
		db: db,
	}
}

func (r *pagesRepository) CreatePages(ctx context.Context, pages entities.Pages) (entities.Pages, error) {
	if err := r.db.Create(&pages).Error; err != nil {
		return entities.Pages{}, err
	}

	return pages, nil
}
