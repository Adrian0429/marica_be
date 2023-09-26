package services

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/repository"
)

type PdfService interface {
}

type pdfService struct {
	pdfRepository repository.PdfRepository
}

func NewPdfService(ir repository.PdfRepository) PdfService {
	return &pdfService{
		pdfRepository: ir,
	}
}

func (is *pdfService) Create(ctx context.Context, pdfCreateDTO dto.PDFCreateDTO) error {
	return nil
}
