package main

import (
	"fmt"

	"github.com/wilzygon/funciones_constructoras/course"
)

func main() {
	Go := &course.Course{
		Name:    "Go desde cero",
		Price:   12.24,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Clases: map[int]string{
			1: "Introducción",
			2: "Estructuras",
			3: "metodos",
		},
	}
	//PrintClases(Go)
	Go.PrintClases()
	Go.ChangePrice(67.12)
	fmt.Println(Go.Price)

}
