package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = `
<html>
<body>
<h1>I Love You! :)</h1> 
</body>
</html>
`

func main() {
	t, _ := template.New("t").Parse(tmpl)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
