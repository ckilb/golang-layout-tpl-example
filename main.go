package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

var tpls map[string]*template.Template

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	pages := []string{
		"home", "about", "404",
	}

	tpls = make(map[string]*template.Template)

	for _, page := range pages {
		tpl, err := template.ParseFiles(
			dir+"/templates/layout.tmpl",
			dir+"/templates/pages/"+page+".tmpl")

		if err != nil {
			panic(err)
		}

		tpls[page] = tpl
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")

		// if no page query param passed, load the home page
		if page == "" {
			page = "home"
		}

		// replace non alpha-numeric characters, for security purposes
		re := regexp.MustCompile("[^a-z0-9_]")
		page = re.ReplaceAllString(page, "")

		// if page doesn't exist, return 404
		if _, ok := tpls[page]; !ok {
			page = "404"
		}

		err := tpls[page].Execute(w, nil)

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":3000", nil)
}
