package helpers

import (
	"fmt"
	"reflect"
)

// GetTypes func .
func GetTypes(data []interface{}) {
	// saber el tipo de dato
	for _, v := range data {
		g := reflect.TypeOf(v)
		fmt.Printf("%v - %v \n", reflect.ValueOf(v), g.Name())
	}
}
