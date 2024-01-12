package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

type (
	WorksheetRepository interface {
		CreateWorksheet(ctx context.Context, worksheet entities.Worksheet) (entities.Worksheet, error)
	}

	worksheetRepository struct {
		db *gorm.DB
	}
)

func NewWorksheetRepository(db *gorm.DB) WorksheetRepository {
	return &worksheetRepository{
		db: db,
	}
}

func (r *worksheetRepository) CreateWorksheet(ctx context.Context, worksheet entities.Worksheet) (entities.Worksheet, error) {
	if err := r.db.Create(&worksheet).Error; err != nil {
		return entities.Worksheet{}, err
	}

	return worksheet, nil
}

func (r *worksheetRepository) DeleteWorksheet(ctx context.Context, worksheet entities.Worksheet) (entities.Worksheet, error) {
	return worksheet, nil
}
