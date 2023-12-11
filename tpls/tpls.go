package tpls

import (
	"embed"
	"errors"
	"html/template"
)

//go:embed **/*.tmpl *.tmpl
var filesystem embed.FS
var tpls map[string]*template.Template

func Init() {
	pages := []string{
		"home", "about",
	}

	tpls = make(map[string]*template.Template)

	funcMap := template.FuncMap{
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, errors.New("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, errors.New("dict keys must be strings")
				}

				dict[key] = values[i+1]
			}
			return dict, nil
		},
	}

	for _, page := range pages {
		tpl := template.New("layout.tmpl").Funcs(funcMap)
		tpl, err := tpl.ParseFS(
			filesystem,
			"layout.tmpl",
			"pages/"+page+".tmpl",
			"sidebar.tmpl",
			"headline.tmpl",
		)

		if err != nil {
			panic(err)
		}

		tpls[page] = tpl
	}
}

func Get(page string) *template.Template {
	return tpls[page]
}
