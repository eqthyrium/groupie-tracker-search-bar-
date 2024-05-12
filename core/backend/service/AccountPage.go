package service

import (
	"net/http"
	"strconv"
	"text/template"

	"Adlet/core/backend/container"
	"Adlet/core/backend/errorness"
)

func AccountPageHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		errorness.ErrorHandler(writer, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("../core/frontend/service/Accountpage.html")
	if err != nil {
		errorness.ErrorHandler(writer, http.StatusInternalServerError)
		return
	}

	ImageIndex := request.FormValue("Image_Index_Value")
	if len(ImageIndex) >= 2 && ImageIndex[0] == '0' {
		errorness.ErrorHandler(writer, http.StatusNotFound)
		return
	}
	result, err := strconv.Atoi(ImageIndex)
	if result < 0 || result >= len(container.Array_Available) || err != nil {
		errorness.ErrorHandler(writer, http.StatusNotFound)
		return
	}
	result = container.Array_Available[result]
	if result >= 52 || result < 0 {
		errorness.ErrorHandler(writer, http.StatusBadRequest)
		return
	}

	artist := container.Mainobject.Artist.Index[result]
	// location := container.Mainobject.Location.Index[result]
	// dates := container.Mainobject.Dates.Index[result]
	relation := container.Mainobject.Relation.Index[result]

	data := struct {
		Image         string
		Name          string
		Members       []string
		CreationDate  int
		FirstAlbum    string
		DatesLocation map[string][]string
	}{
		Image:         artist.Image,
		Name:          artist.Name,
		Members:       artist.Members,
		CreationDate:  artist.CreationDate,
		FirstAlbum:    artist.FirstAlbum,
		DatesLocation: relation.DatesLocations,
	}

	err = tmpl.Execute(writer, data)
	if err != nil {
		errorness.ErrorHandler(writer, http.StatusInternalServerError)
		return
	}
}
