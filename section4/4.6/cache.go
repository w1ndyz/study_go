package main

import (
	"html/template"
	"net/http"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			A string
			B template.HTML
		}
		data.A = "<b>Hello!</b>"
		data.B = "<b>Hello!</b>"

		t.Execute(w, data)
	})
	http.ListenAndServe(":8000", nil)
}
