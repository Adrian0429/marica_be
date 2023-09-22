package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Caknoooo/golang-clean_template/dto"
)

const (
	LOCALHOST  = "http://localhost:8888/api/"
	IMAGE      = "image/get/"
	AUDIO      = "audio/get/"
	PRODUCTION = "http://apibyriski.my.id/api/"
)

func DecodeBase64(base64String string) ([]byte, error) {
	parts := strings.SplitN(base64String, ",", 2)
	if len(parts) != 2 {
		return nil, dto.ErrBase64Format
	}

	base64Data := parts[1]

	decodeBytes, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return nil, dto.ErrDecodeBase64
	}

	return decodeBytes, nil
}

func ToBase64(b []byte) (string, error) {
	encodeBytes := base64.StdEncoding.EncodeToString(b)
	if encodeBytes == "" {
		return "", dto.ErrToBase64
	}
	if encodeBytes == "" {
		return "", dto.ErrNull
	}

	return encodeBytes, nil
}

func IsBase64(file multipart.FileHeader) (string, error) {
	fileData, err := file.Open()
	if err != nil {
		return "", dto.ErrOpenFileMultipart
	}

	defer fileData.Close()

	bytes, err := io.ReadAll(fileData)
	if err != nil {
		return "", dto.ErrOpenIoReader
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/svg+xml":
		base64Encoding += "data:image/svg+xml;base64,"
	case "image/gif": 
		base64Encoding += "data:image/gif;base64,"
	case "application/pdf":
		base64Encoding += "data:application/pdf;base64,"
	case "audio/mpeg":
		base64Encoding += "data:audio/mpeg;base64,"
	case "audio/wave":
		base64Encoding += "data:audio/wave;base64,"
	case "audio/mp3":
		base64Encoding += "data:audio/mp3;base64,"
	case "application/octet-stream":
		base64Encoding += "data:application/octet-stream;base64,"
	default:
		return "", fmt.Errorf("Unsupported file type: %s", mimeType)
	}

	base64, err := ToBase64(bytes)
	if err != nil {
		return "", dto.ErrToBase64
	}
	base64Encoding += base64

	return base64Encoding, nil
}

func SaveImage(base64 string, path string, dirname string, filename string) error {
	data, err := DecodeBase64(base64)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path+"/"+dirname, 0666)
	if err != nil {
		return err
	}

	err = os.WriteFile(path+"/"+dirname+"/"+dirname+"_"+filename, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func SaveAudio(base64 string, path string, dirname string, filename string) error {
	data, err := DecodeBase64(base64)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path+"/"+dirname, 0666)
	if err != nil {
		return err
	}

	err = os.WriteFile(path+"/"+dirname+"/"+dirname+"_"+filename, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func GetImage(dirfile string, filename string) (string, error) {
	file, err := os.Open(dirfile + "/" + filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	base64, err := ToBase64(bytes)
	if err != nil {
		return "", err
	}

	return base64, nil
}

func GenerateFileName(path string, dirname string, filename string) string {
	if os.Getenv("APP_ENV") != "Production" {
		return LOCALHOST + IMAGE + path + "/" + dirname + "/" + dirname + "_" + filename
	}
	return PRODUCTION + IMAGE + path + "/" + dirname + "/" + dirname + "_" + filename
}

func GenerateAudioFileName(path string, dirname string, filename string) string {
	if os.Getenv("APP_ENV") != "Production" {
		return LOCALHOST + AUDIO + path + "/" + dirname + "/" + dirname + "_" + filename
	}
	return PRODUCTION + AUDIO + path + "/" + dirname + "/" + dirname + "_" + filename
}

func Getextension(filename string) string {
	return filepath.Ext(filename)
}
