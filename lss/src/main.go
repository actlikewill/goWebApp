package main

import (
	"database/sql"
	"fileserver/src/controller"
	"fileserver/src/middleware"
	"fileserver/src/model"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "net/http/pprof"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)	
	go http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddlware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {	
	db, err := sql.Open("postgres", "postgres://lss:lss@localhost:5433/lss?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory")
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read the content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "'")
		}
		result[fi.Name()] = tmpl
	}
	return result
}