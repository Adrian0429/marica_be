package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type AudioService interface {
}

type audioService struct {
	audioRepository repository.AudioRepository
}

func NewAudioService(ir repository.AudioRepository) AudioService {
	return &audioService{
		audioRepository: ir,
	}
}

func (is *audioService) Create(ctx context.Context, audioCreateDTO dto.AudioCreateDTO) error {
	return nil
}
