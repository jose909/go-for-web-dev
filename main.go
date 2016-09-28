
package main

import (
	"fmt"
	"net/http"
	"html/template"

	_ "github.com/mattn/go-sqlite3"
  "database/sql"
)

type Page struct {
  Name string
  DBStatus bool
  }

func main() {
	template := template.Must(template.ParseFiles("templates/index.html"))

	db, _ := sql.Open("sqlite3", "dev.db")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		p := Page{Name: "Jose"}
		if name := r.FormValue("name"); name != ""{
			p.Name = name
		}

		p.DBStatus = db.Ping() == nil

		if err := template.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

	fmt.Println(http.ListenAndServe(":8080", nil))
}