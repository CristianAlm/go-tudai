/**El trabajo practico consta de hacer una función que dado una cadena de caracteres devuelva una estructura*/

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

var primera string
var segunda string
var terceraInt int
var terceraString string
var cadena string = "NN040001"

func main() {

	for range cadena {
		primera = cadena[0:2]
		segunda = cadena[2:4]
		terceraString = cadena[4:]
	}

	fmt.Println(primera)
	fmt.Println(segunda)
	fmt.Println(terceraString)
	fmt.Println(terceraInt)

	terceraInt, err := strconv.Atoi(terceraString)
	fmt.Println(terceraInt, err, reflect.TypeOf(terceraInt))

	p := Result{primera, segunda, terceraInt}
	fmt.Println(p)

}
