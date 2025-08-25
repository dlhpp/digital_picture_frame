package internal

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func GetImageStore() *ImageStore {
	// Initialize image store
	dataDir := "./data" // Change to your data directory path
	images, err := scanImages(dataDir)
	if err != nil {
		panic("getImageStore: Failed to scan images: " + err.Error())
	}
	slog.Info("getImageStore:", "len(images)", len(images))

	// DLH:  We never said "new ImageStore" - we just created a pointer to an ImageStore struct.
	// DLH:  I just realized, we're actually creating the ImageStore struct here with the curly braces.
	store := &ImageStore{Images: images, ImageSubscript: 0}
	return store
}

// scanImages recursively scans the directory for image files
func scanImages(dataDir string) ([]string, error) {
	var images []string
	err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Check for common image extensions
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
				absPath, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				images = append(images, absPath)
			}
		}
		return nil
	})
	return images, err
}
