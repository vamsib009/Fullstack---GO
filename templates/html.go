package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	fmt.Println("server started")
	http.Handle("/", new(vamsi))
	http.ListenAndServe(":8080", nil)

}
