package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func SaveUploadedFile(file multipart.File, fileName string) (string, error) {

	validExtensions := []string{".jpg", ".jpeg", ".png"}

	ext := filepath.Ext(fileName)
	ext = strings.ToLower(ext)

	if !isValidExtension(ext, validExtensions) {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	uploadDir := "uploads"

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	validFileName := strings.ReplaceAll(fileName, " ", "_")
	filePath := filepath.Join(uploadDir, validFileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func isValidExtension(ext string, validExtensions []string) bool {
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
