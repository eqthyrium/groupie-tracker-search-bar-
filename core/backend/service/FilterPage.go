package service

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"Adlet/core/backend/container"

	"Adlet/core/backend/errorness"
	"Adlet/core/backend/service/support"
)

var (
	PassedValue string
	TypeValue   string
)

type Data struct {
	ImageValue1 []string
}

var data Data

func FilterPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// fmt.Println("Method not Allowed error in the FilterPage function")
		errorness.ErrorHandler(w, http.StatusMethodNotAllowed)
	}

	// fmt.Println(r.FormValue("CareerYear"))
	// fmt.Println(len(r.FormValue("Album")))
	// fmt.Println(len(r.FormValue("LocationValue")))
	// fmt.Println(len(r.FormValue("Members1")))

	PassedValue := r.FormValue("HashTableValue")
	PassedValue = strings.TrimSpace(PassedValue)
	fmt.Println("Passed Value From the Front is:", PassedValue)

	var tmpl *template.Template
	var err error
	var name []string

	if !support.CheckValidation(w, PassedValue) && len(PassedValue) != 0 {
		fmt.Println("There is such supported Data in the Database, error in the Filter Function")
		tmpl, err = template.ParseFiles("../core/frontend/errorness/Nonexistence.html")
		if err != nil {
			fmt.Println("Internal Server Error From thr template to Nonexistence file in FilterPage function")
			errorness.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	} else {
		for i := 0; i < len(container.Array_Available); i++ {
			name = append(name, container.Mainobject.Artist.Index[container.Array_Available[i]].Image)
		}
		data.ImageValue1 = name

		tmpl, err = template.ParseFiles("../core/frontend/service/Choosepage.html")
		if err != nil {
			fmt.Println("Internal Server Error From thr template to Choosepage file in Filterpage Function ")
			errorness.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println("Execution Error for template in Filterpage fucntion")
		errorness.ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	// tmpl, err = template.ParseFiles("../core/frontend/service/Filterpage.html") // XOR account XOR Error page with link to main page
	// if err != nil {
	// 	errorness.ErrorHandler(w, http.StatusInternalServerError)
	// 	return
	// }
}
