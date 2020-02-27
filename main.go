package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "port")

var tmpl = `
<html>
<body>
	<h1>I Love You, Mahal! 😍 💝 😍</h1>
</body>
</html>
`

func main() {
	flag.Parse()
	t, _ := template.New("t").Parse(tmpl)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
