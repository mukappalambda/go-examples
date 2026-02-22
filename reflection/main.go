package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID    int
	Name  string
	Score float32
	Tasks []string
	Done  bool
}

func main() {
	user := User{
		ID:    1,
		Name:  "alex",
		Score: 10.1,
		Tasks: []string{"job1", "job2"},
		Done:  true,
	}

	t := reflect.TypeFor[User]()
	v := reflect.ValueOf(user)
	k := t.Kind()

	fmt.Println(t)
	fmt.Println(t.NumField(), t.Field(3).Type)
	fmt.Println(v)
	fmt.Println(k)

	a := 1
	tint := reflect.TypeFor[int]()
	vint := reflect.ValueOf(a)
	fmt.Println(tint, vint)

	names := []string{"alex", "bob", "mark"}
	tnames := reflect.TypeFor[[]string]()
	vnames := reflect.ValueOf(names)
	fmt.Println(tnames, vnames)

	myMap := map[string]string{
		"name":   "alex",
		"email":  "alex@gmail.com",
		"gender": "male",
	}

	tMyMap := reflect.TypeFor[map[string]string]()
	vMyMap := reflect.ValueOf(myMap)

	fmt.Println(tMyMap, vMyMap, tMyMap.Elem())

	c := make(chan string)
	tc := reflect.TypeFor[chan string]()
	vc := reflect.ValueOf(c)
	fmt.Println(tc, vc)
}
