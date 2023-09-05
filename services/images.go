package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type ImageService interface {
}

type imageService struct {
	imageRepository repository.ImageRepository
}

func NewImageService(ir repository.ImageRepository) ImageService {
	return &imageService{
		imageRepository: ir,
	}
}

func (is *imageService) Create(ctx context.Context, imageCreateDTO dto.ImageCreateDTO) error {
	return nil
}
