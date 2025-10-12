package internal

import (
	"html/template"
	"os"

	"github.com/dlhpp/digital_picture_frame/logging"
)

func getTemplate() *template.Template {
	logging.Log("getTemplate", 5, "entering")
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		logging.Log("getTemplate", 9, "error getting template: "+err.Error())
		os.Exit(1)
	}
	return tmpl
}
