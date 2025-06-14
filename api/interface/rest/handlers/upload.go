package handlers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ggicci/httpin"
)

type UploadHandler struct {
	UploadDir   string
	TemplateDir string
}

func (h *UploadHandler) UploadTemplate(input *httpin.File, slug string) error {
	// 50MB max size
	if input.Size() > 50*1024*1024 {
		return fmt.Errorf("template too large: %d bytes", input.Size())
	}

	file, err := input.OpenReceiveStream()
	if err != nil {
		return fmt.Errorf("failed to open image file: %w", err)
	}

	// 3. Validate image type
	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		return fmt.Errorf("failed to read image: %w", err)
	}

	filetype := http.DetectContentType(buff)
	if filetype != "application/zip" {
		return fmt.Errorf("invalid zip type: %s", filetype)
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to reset file pointer: %w", err)
	}

	// 4. Create upload directory if not exists
	if err := os.MkdirAll(h.TemplateDir, 0o755); err != nil {
		return fmt.Errorf("failed to create upload directory: %w", err)
	}

	// 5. Unzip the file to the target directory
	zipReader, err := zip.NewReader(file, input.Size())
	if err != nil {
		return fmt.Errorf("failed to read zip file: %w", err)
	}

	// Create a subdirectory for this template
	templateExtractPath := filepath.Join(h.TemplateDir, slug)
	if err := os.MkdirAll(templateExtractPath, 0o755); err != nil {
		return fmt.Errorf("failed to create template directory: %w", err)
	}

	// Extract each file from the zip archive
	for _, zipFile := range zipReader.File {
		// Get the absolute path for security check
		fullPath := filepath.Join(templateExtractPath, zipFile.Name)

		// Check for ZipSlip vulnerability
		if !strings.HasPrefix(fullPath, filepath.Clean(templateExtractPath)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", zipFile.Name)
		}

		if zipFile.FileInfo().IsDir() {
			// Create directory
			if err := os.MkdirAll(fullPath, zipFile.Mode()); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		// Create parent directories if needed
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			return fmt.Errorf("failed to create parent directory: %w", err)
		}

		// Open the file inside the zip archive
		srcFile, err := zipFile.Open()
		if err != nil {
			return fmt.Errorf("failed to open zipped file: %w", err)
		}
		defer srcFile.Close()

		// Create the destination file
		dstFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
		if err != nil {
			return fmt.Errorf("failed to create destination file: %w", err)
		}
		defer dstFile.Close()

		// Copy the file contents
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file contents: %w", err)
		}
	}
	return nil
}

func (h *UploadHandler) UploadImage(input *httpin.File) (string, error) {
	// 5MB max size
	// 1. Parse multipart form (5MB max for images)
	if input.Size() > 5*1024*1024 {
		return "", fmt.Errorf("image too large: %d bytes", input.Size())
	}

	file, err := input.OpenReceiveStream()
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}

	// 3. Validate image type
	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		return "", fmt.Errorf("failed to read image: %w", err)
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		return "", fmt.Errorf("invalid image type: %s", filetype)
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to reset file pointer: %w", err)
	}

	// 4. Create upload directory if not exists
	if err := os.MkdirAll(h.UploadDir, 0o755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// 5. Generate unique filename
	ext := filepath.Ext(input.Filename())
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(h.UploadDir, newFilename)

	// 6. Save file
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		return "", fmt.Errorf("failed to write image: %w", err)
	}

	// 7. Return public URL
	publicURL := "/public/" + newFilename
	return publicURL, nil
}
