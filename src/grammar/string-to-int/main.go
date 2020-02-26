package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("1")
	fmt.Println(i, err)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(i).Kind())
}
