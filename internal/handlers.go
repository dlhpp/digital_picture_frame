package internal

import (
	"html/template"
	"net/http"

	"github.com/dlhpp/digital_picture_frame/logging"
)

var tmpl *template.Template

func SetupHttpHandlers(store *ImageStore) {
	// Set up HTTP handlers
	logging.Log("SetupHttpHandlers", 5, "entering")

	tmpl = getTemplate()

	// These register the functions, they DO NOT call the functions.
	http.HandleFunc("/", store.indexHandler)
	http.HandleFunc("/next", store.nextImageHandler)
}

// indexHandler serves the HTML template
func (store *ImageStore) indexHandler(w http.ResponseWriter, r *http.Request) {

	logging.Log("indexHandler", 5, "entering - will return main parent HTML page")

	// Data for template (optional substitutions)
	data := struct {
		Title    string
		Fadetime int
		Holdtime int
	}{
		Title:    store.Title,
		Fadetime: store.Fadetime,
		Holdtime: store.Holdtime,
	}

	// Execute template
	if err := tmpl.Execute(w, data); err != nil {
		logging.Log("indexHandler", 9, "error executing main html index template: "+err.Error())
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

// nextImageHandler serves a random image
// DLH:  This is a method on the ImageStore struct.  Seems identical to object oriented programming.
func (store *ImageStore) nextImageHandler(w http.ResponseWriter, r *http.Request) {
	if len(store.Images) == 0 {
		http.Error(w, "No images found", http.StatusNotFound)
		return
	}

	// Select a random image
	// rand.Seed(time.Now().UnixNano())
	// imagePath := store.Images[rand.Intn(len(store.Images))]
	imagePath := store.Images[store.ImageSubscript]
	logging.Log("nextImageHandler", 3, "subscript", store.ImageSubscript, "imagePath", imagePath)
	store.ImageSubscript++
	if store.ImageSubscript >= len(store.Images) {
		store.ImageSubscript = 0
	}

	// Serve the image file
	http.ServeFile(w, r, imagePath)
}
