package dto

type (
	AudioCreateDTO struct {
		Path       string `json:"path" form:"path"`
		FileName   string `json:"file_name" form:"file_name"`
		Base64Data string `json:"base64Data" form:"base64Data"`
	}

	AudioUploadDTO struct {
		AudioForm string `json:"image_form" form:"image_form"`
	}
)
