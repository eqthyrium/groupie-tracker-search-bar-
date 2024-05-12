package service

import (
	"html/template"
	"net/http"

	"Adlet/core/backend/container"
	"Adlet/core/backend/errorness"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorness.ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorness.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("../core/frontend/service/Filterpage.html")
	if err != nil {
		errorness.ErrorHandler(w, http.StatusInternalServerError)
	}

	// fmt.Println(r.FormValue("CareerYear"))
	// fmt.Println(r.FormValue("Members1"))
	// fmt.Println(r.FormValue("Album"))

	data := struct { // For search bar
		Hashtable map[string]map[string][]int
	}{
		Hashtable: container.MapAll,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		errorness.ErrorHandler(w, http.StatusInternalServerError)
	}
}
