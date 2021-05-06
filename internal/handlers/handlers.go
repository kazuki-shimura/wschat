package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// var views = jet.NewSet(
// 	jet.NewOSFileSystemLoader("./html"),
// 	jet.InDevelopmentMode(),
// )

// func Home(w http.ResponseWriter, r *http.Request) {
// 	err := renderPage(w, "home.jet", nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
// 	view, err := views.GetTemplate(tmpl)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	err = view.Execute(w, data, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	return nil
// }

func top(w http.ResponseWriter, r *http.Request) {

}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	http.HandleFunc("/", top)

	return http.ListenAndServe(":8080", nil)
}
