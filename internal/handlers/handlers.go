package handlers

import (
	"html/template"
	"net/http"

	"github.com/maxzhirnov/utm-builder/internal/dictionary"
	"github.com/maxzhirnov/utm-builder/internal/utm"
)

var templates = template.Must(template.ParseGlob("internal/templates/*.html"))

func FormHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Sources []string
		Mediums []string
	}{
		Sources: dictionary.Sources(),
		Mediums: dictionary.Mediums(),
	}
	templates.ExecuteTemplate(w, "base.html", data)
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	source := r.FormValue("utm_source")
	if source == "" {
		source = r.FormValue("utm_source_custom")
	}
	medium := r.FormValue("utm_medium")
	if medium == "" {
		medium = r.FormValue("utm_medium_custom")
	}
	params := utm.UTMParams{
		URL:      r.FormValue("url"),
		Source:   source,
		Medium:   medium,
		Term:     r.FormValue("utm_term"),
		Content:  r.FormValue("utm_content"),
		Campaign: r.FormValue("utm_campaign"),
	}
	link := params.BuildLink()
	data := struct{ Link string }{Link: link}
	templates.ExecuteTemplate(w, "result", data)
}

// GetOptionsHandler returns the current UTM options as HTML
func GetOptionsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the type of options requested (sources or mediums)
	optionType := r.URL.Query().Get("type")
	if optionType != "sources" && optionType != "mediums" {
		http.Error(w, "Invalid option type", http.StatusBadRequest)
		return
	}

	// Get the current options
	var options []string
	if optionType == "sources" {
		options = dictionary.Sources()
	} else {
		options = dictionary.Mediums()
	}

	// Render the options as HTML
	tmpl := template.Must(template.New("options").Parse(`
		<option value="">Select {{.Type}}</option>
		{{range .Options}}
		<option value="{{.}}">{{.}}</option>
		{{end}}
		<option value="custom">Custom {{.Type}}</option>
	`))

	data := struct {
		Type    string
		Options []string
	}{
		Type:    optionType,
		Options: options,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering options", http.StatusInternalServerError)
		return
	}
}
