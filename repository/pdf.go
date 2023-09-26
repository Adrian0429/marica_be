package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Caknoooo/golang-clean_template/entities"
)

type PdfRepository interface {
}

type pdfRepository struct {
	db *gorm.DB
}

func NewPdfRepository(db *gorm.DB) *pdfRepository {
	return &pdfRepository{
		db: db,
	}
}

func (r *pdfRepository) Create(ctx context.Context, pdf entities.PDF) (entities.PDF, error) {
	if err := r.db.WithContext(ctx).Create(&pdf).Error; err != nil {
		return entities.PDF{}, err
	}

	return pdf, nil
}
