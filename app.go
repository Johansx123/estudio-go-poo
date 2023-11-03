package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c Course) PrintClasses() {
	text := "Las Clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func main() {
	Go := Course{
		Name:    "Go desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "introduccion",
			2: "Estructuras",
			3: "maps",
		},
	}

	Css := Course{Name: "Css desde cero", IsFree: true}
	js := Course{}
	js.Name = "Curso JS"
	js.UserIDs = []uint{12, 67}

	fmt.Printf("%+v\n", Css)

	Go.PrintClasses()
}
