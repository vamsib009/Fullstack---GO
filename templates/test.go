package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

type data struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("land.html"))
}

func main() {

	res, err := http.Get("http://careers-data.benzinga.com/rest/richquoteDelayed?symbols=GE")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	var s = new(Jsonresponce)

	err = json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}
	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatal(err)
	}

}
