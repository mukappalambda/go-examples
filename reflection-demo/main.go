package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id    int
	Name  string
	Score float32
	Tasks []string
	Done  bool
}

func main() {
	user := User{
		Id:    1,
		Name:  "alex",
		Score: 10.1,
		Tasks: []string{"job1", "job2"},
		Done:  true,
	}

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)
	k := t.Kind()

	fmt.Println(t)
	fmt.Println(t.NumField(), t.Field(3).Type)
	fmt.Println(v)
	fmt.Println(k)

	a := 1
	tint := reflect.TypeOf(a)
	vint := reflect.ValueOf(a)
	fmt.Println(tint, vint)

	names := []string{"alex", "bob", "mark"}
	tnames := reflect.TypeOf(names)
	vnames := reflect.ValueOf(names)
	fmt.Println(tnames, vnames)

	myMap := map[string]string{
		"name":   "alex",
		"email":  "alex@gmail.com",
		"gender": "male",
	}

	tMyMap := reflect.TypeOf(myMap)
	vMyMap := reflect.ValueOf(myMap)

	fmt.Println(tMyMap, vMyMap, tMyMap.Elem())

	c := make(chan string)
	tc := reflect.TypeOf(c)
	vc := reflect.ValueOf(c)
	fmt.Println(tc, vc)

}
