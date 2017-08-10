package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "stockinstance.citsm1ymtilj.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "stock"
	password = "Vamsi1994"
	dbname   = "postgres"
)

type vamsi struct {
}

func (vamsi *vamsi) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]

	fmt.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "image/jpg"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("content Type", contentType)
		w.Write(data)
	} else {

		w.Write([]byte("404 vamsi - " + http.StatusText(404)))
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("server started")
	http.Handle("/", new(vamsi))
	http.ListenAndServe(":8080", nil)

}
