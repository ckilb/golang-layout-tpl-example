package main

import (
	"ckilb/golang-layout-tpl/tpls"
	"net/http"
)

func main() {
	// get the current working directory
	tpls.Init()

	// create & start a web server that will render the templates
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpls.Get("home").Execute(w, map[string]interface{}{
			"name": "Fred",
		})

		if err != nil {
			panic(err)
		}
	})

	// create & start a web server that will render the template
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		err := tpls.Get("about").Execute(w, nil)

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":3210", nil)
}
