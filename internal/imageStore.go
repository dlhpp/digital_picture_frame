package internal

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func ShuffleImages(store *ImageStore) {
	// rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(store.Images), func(i, j int) {
		store.Images[i], store.Images[j] = store.Images[j], store.Images[i]
	})
}

func GetImageStore(flags *FlagSettings) *ImageStore {
	// Initialize image store
	dataDir := "./data" // Change to your data directory path
	images, err := scanImages(dataDir)
	if err != nil {
		panic("getImageStore: Failed to scan images: " + err.Error())
	}
	// slog.Info("getImageStore:", "len(images)", len(images))
	DLHLog("GetImageStore", 5, fmt.Sprintf("len(images) = %d", len(images)))

	// DLH:  We never said "new ImageStore" - we just created a pointer to an ImageStore struct.
	// DLH:  I just realized, we're actually creating the ImageStore struct here with the curly braces.
	store := &ImageStore{Images: images, ImageSubscript: 0}

	if flags.Random {
		ShuffleImages(store)
		DLHLog("GetImageStore", 5, "Shuffled images for random display.")
	}

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
