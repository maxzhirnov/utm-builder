package handlers

import (
    "html/template"
    "net/http"
    "github.com/maxzhirnov/utm-builder/internal/utm"
    "github.com/maxzhirnov/utm-builder/internal/dictionary"
)

var templates = template.Must(template.ParseGlob("internal/templates/*.html"))

func FormHandler(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Sources []string
        Mediums []string
    }{
        Sources: dictionary.Sources,
        Mediums: dictionary.Mediums,
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
