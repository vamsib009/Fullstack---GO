package main

import (
	"fmt"
	"reflect"
)

type CourseAssignment struct {
	Semester int `json:"semester"  xml:"semester"`
}

func main() {
	ca := CourseAssignment{}
	st := reflect.TypeOf(ca)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("json"))
}
