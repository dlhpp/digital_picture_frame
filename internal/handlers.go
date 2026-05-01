package internal

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/dlhpp/digital_picture_frame/logging"
)

var tmpl *template.Template

func createStaticHandle(url string, filePath string, useExactPath bool) {
	logging.Log("createStaticHandle", 5, fmt.Sprintf("entering: url = %s, filePath = %s", url, filePath))

	fileHandler := http.FileServer(http.Dir(filePath))
	urlHandler := http.StripPrefix(url, fileHandler)

	handler := urlHandler
	if useExactPath {
		handler = fileHandler
	}

	// Wrap the handler to access the request in order to get the url path.
	wrappedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Log("createStaticHandle", 5, fmt.Sprintf("RESULT: useExactPath = %t, url path = %s, filePath = %s", useExactPath, r.URL.Path, filePath))
		handler.ServeHTTP(w, r) // Delegate to the original handler
	})

	http.Handle(url, wrappedHandler)
}

func SetupHttpHandlers(store *ImageStore) {
	// Set up HTTP handlers
	logging.Log("SetupHttpHandlers", 5, "entering")

	tmpl = getTemplate()

	http.HandleFunc("/", store.indexHandler)

	http.HandleFunc("/next", store.nextImageHandler)

	createStaticHandle("/static/", "static", false)

	createStaticHandle("/favicon.ico", "./static/icons/favicon_fandom.ico", true)

	// 2026/04/29 17:12:39 [05/05] SetupHttpHandlers: path = /favicon.ico, will return favicon_fandom.ico
	// http.HandleFunc("/favicon.ico", func(response http.ResponseWriter, request *http.Request) {
	// 	logging.Log("SetupHttpHandlers", 5, fmt.Sprintf("path = %s, will return favicon_fandom.ico", request.URL.Path))
	// 	http.ServeFile(response, request, "favicon_fandom.ico")
	// })
}

// indexHandler serves the HTML template
func (store *ImageStore) indexHandler(w http.ResponseWriter, r *http.Request) {

	// logging.Log("indexHandler", 5, fmt.Sprintf("Handling request for path: %s\n", r.URL.Path))
	logging.Log("indexHandler", 5, fmt.Sprintf("entering: path = %s, will return main parent HTML page", r.URL.Path))

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
	// logging.Log("nextImageHandler", 3, "subscript", store.ImageSubscript, "imagePath", imagePath)
	logging.Log("nextImageHandler", 3, fmt.Sprintf("path = %s, subscript = %d, imagePath = %s", r.URL.Path, store.ImageSubscript, imagePath))
	store.ImageSubscript++
	if store.ImageSubscript >= len(store.Images) {
		store.ImageSubscript = 0
	}

	// Serve the image file
	http.ServeFile(w, r, imagePath)
}
