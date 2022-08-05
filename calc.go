package main

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

type ContactDetails struct {
	float1    string
	float2    string
	Operation string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		var details = ContactDetails{
			float1:    r.FormValue("input1"),
			float2:    r.FormValue("input2"),
			Operation: r.FormValue("ope"),
		} // do something with details

		float1, _ := strconv.ParseFloat(details.float1, 64)
		float2, _ := strconv.ParseFloat(details.float2, 64)

		var res float64

		if details.Operation == "Add" {
			res = float1 + float2
		} else if details.Operation == "Subtract" {
			res = float1 - float2
		} else if details.Operation == "Divide" {
			res = float1 / float2
		} else if details.Operation == "Multiply" {
			res = float1 * float2
		}

		res = math.Round(res*100) / 100

		_ = details

		tmpl.Execute(w, struct{ Res float64 }{res})

	})
	http.ListenAndServe(":8080", nil)

}
