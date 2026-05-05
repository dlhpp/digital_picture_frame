package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/dlhpp/digital_picture_frame/logging"
)

var tmpl *template.Template

// func createStaticHandle(url string, filePath string) {
// 	logging.Log("createStaticHandle", 5, fmt.Sprintf("entering: url = %s, filePath = %s", url, filePath))
// 	fileHandler := http.FileServer(http.Dir(filePath))
// 	urlHandler := http.StripPrefix(url, fileHandler)
// 	// Wrap the handler to access the request in order to get the url path.
// 	wrappedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		logging.Log("createStaticHandle", 5, fmt.Sprintf("RESULT: url path = %s, filePath = %s", r.URL.Path, filePath))
// 		urlHandler.ServeHTTP(w, r) // Delegate to the original handler
// 	})
// 	http.Handle(url, wrappedHandler)
// }

func (store *ImageStore) rootHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	switch {
	case path == "/", path == "/index.html":
		logging.Log("rootHandler", 5, fmt.Sprintf("path = %s, calling indexHandler", r.URL.Path))
		store.indexHandler(w, r)

	case path == "/next":
		logging.Log("rootHandler", 5, fmt.Sprintf("path = %s, calling nextImageHandler", r.URL.Path))
		store.nextImageHandler(w, r)

	case path == "/favicon.ico":
		logging.Log("rootHandler", 5, fmt.Sprintf("path = %s, returning favicon.ico", r.URL.Path))
		http.ServeFile(w, r, "static/icons/favicon_fandom.ico")

	case strings.HasPrefix(path, "/static/") && len(path) > len("/static/"):
		// check len to prevent returning a directory listing if /static/ is requested without a file
		logging.Log("rootHandler", 5, fmt.Sprintf("path = %s, returning favicon.ico", r.URL.Path))
		http.ServeFile(w, r, strings.TrimPrefix(path, "/"))

	default:
		logging.Log("rootHandler", 5, fmt.Sprintf("path = %s, PATH NOT FOUND", r.URL.Path))
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func SetupHttpHandlers(store *ImageStore) {
	// Set up HTTP handlers
	logging.Log("SetupHttpHandlers", 5, "entering")

	tmpl = getTemplate()

	// http.HandleFunc("/", store.indexHandler)
	http.HandleFunc("/", store.rootHandler)

	// http.HandleFunc("/next", store.nextImageHandler)

	// http.HandleFunc("/favicon.ico", faviconHandler)

	// createStaticHandle("/static/", "static")
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
