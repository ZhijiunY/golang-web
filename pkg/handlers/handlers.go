package handlers

import (
	"net/http"

	"github.com/ZhijiunY/golang-web/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// call renderTemplate
	render.RenderTemplate(w, "home.page.tmpl")

}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
