package internal

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/dlhpp/digital_picture_frame/logging"
	"github.com/dlhpp/digital_picture_frame/yaml"
)

func ShuffleImages(store *ImageStore) {
	// rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(store.Images), func(i, j int) {
		store.Images[i], store.Images[j] = store.Images[j], store.Images[i]
	})
}

func GetImageStore(yamlConfig *map[string]any, flags *FlagSettings) *ImageStore {
	// Initialize image store
	// dataDir := yaml.GetString(yamlConfig, "picture-directories.0", "--not-found--") // "./data" // Change to your data directory path
	dataDir := yaml.GetStringArray(yamlConfig, "picture-directories", []string{"--not-found--"}) // "./data" // Change to your data directory path
	var imageStore []string

	for idx, folder := range dataDir {

		logging.Log("GetImageStore", 5, fmt.Sprintf("folder[%02d] = %s", idx, folder))
		err := scanImages(folder, &imageStore)
		if err != nil {
			panic("getImageStore: Failed to scan images: " + err.Error())
		}
	}

	// slog.Info("getImageStore:", "len(images)", len(images))
	logging.Log("GetImageStore", 5, fmt.Sprintf("len(imageStore) = %d", len(imageStore)))

	// DLH:  We never said "new ImageStore" - we just created a pointer to an ImageStore struct.
	// DLH:  I just realized, we're actually creating the ImageStore struct here with the curly braces.
	store := &ImageStore{Images: imageStore, ImageSubscript: 0}

	if flags.Random {
		ShuffleImages(store)
		logging.Log("GetImageStore", 5, "Shuffled images for random display.")
	}

	return store
}

// scanImages recursively scans the directory for image files
func scanImages(dataDir string, imageStore *[]string) error {
	logging.Log("scanImages", 5, fmt.Sprintf("dataDir = %s", dataDir))
	err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			logging.Log("scanImages", 5, fmt.Sprintf("dataDir = %s, path = %s", dataDir, path))
			// Check for common image extensions
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" {
				absPath, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				*imageStore = append(*imageStore, absPath)
			}
		}
		return nil
	})

	return err
}
