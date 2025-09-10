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

func GetImageStore(yamlConfig *map[string]any) *ImageStore {
	// Initialize image store
	// dataDir := yaml.GetString(yamlConfig, "picture-directories.0", "--not-found--")
	dataDir := yaml.GetStringArray(yamlConfig, "picture-directories", []string{"--not-found--"})
	title := yaml.GetString(yamlConfig, "title", "DLH Slideshow") // time in seconds to transition from one pic to the next
	fadetime := yaml.GetInt(yamlConfig, "fadetime", 3)            // time in seconds to transition from one pic to the next
	holdtime := yaml.GetInt(yamlConfig, "holdtime", 20000)        // time in milliseconds to display one pic
	randomize := (yaml.Get(yamlConfig, "random", true)).(bool)
	logging.Log("GetImageStore", 5, fmt.Sprintf("randomize=%t, fadetime=%d, holdtime=%d, len(dataDir)=%d, title=%s", randomize, fadetime, holdtime, len(dataDir), title))

	var imageStore []string

	for idx, folder := range dataDir {
		logging.Log("GetImageStore", 5, fmt.Sprintf("folder[%02d] = %s", idx, folder))
		err := scanImages(folder, &imageStore)
		if err != nil {
			panic("getImageStore: Failed to scan images: " + err.Error())
		}
	}

	logging.Log("GetImageStore", 5, fmt.Sprintf("len(imageStore) = %d", len(imageStore)))

	store := &ImageStore{Images: imageStore, ImageSubscript: 0, Fadetime: fadetime, Holdtime: holdtime, Title: title}

	if randomize {
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
			logging.Log("scanImages", 1, fmt.Sprintf("dataDir = %s, path = %s", dataDir, path))
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
