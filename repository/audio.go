package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Caknoooo/golang-clean_template/entities"
)

type AudioRepository interface {
}

type audioRepository struct {
	db *gorm.DB
}

func NewAudioRepository(db *gorm.DB) *audioRepository {
	return &audioRepository{
		db: db,
	}
}

func (r *audioRepository) Create(ctx context.Context, audio entities.Audio) (entities.Audio, error) {
	if err := r.db.WithContext(ctx).Create(&audio).Error; err != nil {
		return entities.Audio{}, err
	}

	return audio, nil
}
