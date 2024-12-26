package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

// `ServeTemplate` takes in a template and appropriate HTTP
// tools, and serves `t` as an HTML string in `w`.
//
// What a handy little function.
func ServeTemplate(t string, w http.ResponseWriter, r *http.Request, c any) {
	// Parse `t` as an HTML template
	tmpl, err := template.New("tmpl").Parse(t)
	if err != nil {
		es := fmt.Sprintf("Error loading template [%v]", err)
		http.Error(w, es, http.StatusInternalServerError)
		fmt.Printf("%s", es)
	}

	// Write the HTML to `w`
	//
	// `qc` is our dynamic content, and `go` is smart enough
	// to do the subtitutions for us.
	err = tmpl.Execute(w, c)
	if err != nil {
		es := fmt.Sprintf("Error executing template [%v]", err)
		http.Error(w, es, http.StatusInternalServerError)
		fmt.Printf("%s", es)
	}
}
