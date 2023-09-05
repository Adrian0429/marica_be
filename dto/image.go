package dto

import (
	"errors"
)

var (
	ErrOpenFileMultipart = errors.New("failed to open file multipart")
	ErrOpenIoReader      = errors.New("failed to open io reader")
	ErrToBase64          = errors.New("failed to convert to base64")
	ErrDecodeBase64      = errors.New("failed to decode base64")
	ErrBase64Format      = errors.New("base64 format not valid")
)

type (
	ImageCreateDTO struct {
		Path       string `json:"path" form:"path"`
		FileName   string `json:"file_name" form:"file_name"`
		Base64Data string `json:"base64Data" form:"base64Data"`
	}

	ImageUploadDTO struct {
		ImageForm string `json:"image_form" form:"image_form"`
	}
)
