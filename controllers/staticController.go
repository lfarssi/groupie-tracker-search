package controllers

import (
	"net/http"
	"os"
)

func StaticController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		folders, err := os.ReadDir("static")
		if err != nil {
			ErrorController(w, r, http.StatusInternalServerError)
			return
		}
		for _, folder := range folders {
			files, _ := os.ReadDir("static/" + folder.Name())
			for _, file := range files {
				if r.URL.Path == "/static/style/"+file.Name() || r.URL.Path == "/static/js/"+file.Name() || r.URL.Path == "/static/images/"+file.Name() {
					fs := http.Dir("static")
					http.StripPrefix("/static/", http.FileServer(fs)).ServeHTTP(w, r)
					return
				}
			}
		}
		ErrorController(w, r, http.StatusNotFound)
		return
	}
	ErrorController(w, r, http.StatusMethodNotAllowed)

}