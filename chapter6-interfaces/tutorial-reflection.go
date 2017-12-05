package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string "namestr"
	age  int
}

func showTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag
		fmt.Println(tag)
	default:
	}
}

func show(i interface{}) {
	switch t := i.(type) {
	case *Person:
		tp := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		tag := tp.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
		fmt.Println(tp)   // *main.Person
		fmt.Println(v)    // &{Ryan 18}
		fmt.Println(tag)  //namestr
		fmt.Println(name) // Ryan
		fmt.Println(t)    // &{Ryan 18}
	}
}

func main() {
	p := Person{"Ryan", 18}
	show(&p)
}
