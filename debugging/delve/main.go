package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Data []float64
}

func main() {
	var p Person
	p.Name = "alex"
	p.Age = 20
	p.Data = []float64{1.1, 2.2, 3.3}
	fmt.Println(p)
	p.appendData([]float64{4.4, 5.5})
	fmt.Println(p)
}

func (pp *Person) appendData(s []float64) {
	p := *pp
	p.Data = append(p.Data, s...)
	*pp = p
}
