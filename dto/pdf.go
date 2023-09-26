package dto

type (
	PDFCreateDTO struct {
		Path       string `json:"path" form:"path"`
		FileName   string `json:"file_name" form:"file_name"`
		Base64Data string `json:"base64Data" form:"base64Data"`
	}

	PDFUploadDTO struct {
		PDFForm string `json:"image_form" form:"image_form"`
	}
)
