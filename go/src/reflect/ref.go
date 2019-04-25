package main

import "fmt"

import "reflect"

type T *T

func f(n int) int {
	var a T
	b := &a

	art := reflect.TypeOf(a)
	arv := reflect.ValueOf(a)
	art2 := arv.Type()
	ark := arv.Kind()

	brt := reflect.TypeOf(b)
	brv := reflect.ValueOf(b)
	brt2 := brv.Type()
	brk := brv.Kind()

	frt := reflect.TypeOf(main)
	frv := reflect.ValueOf(main)
	frt2 := brv.Type()
	frk := brv.Kind()

	fmt.Println()
	fmt.Println("a: ", a)
	fmt.Printf("art: %T\n", art)
	fmt.Println("  art: ", art)
	fmt.Printf("arv: %v\n", arv)
	fmt.Println("art2: ", art2)
	fmt.Println("ark: ", ark)

	fmt.Println()
	fmt.Println("b: ", b)
	fmt.Printf("brt: %T\n", brt)
	fmt.Println("  brt: ", brt)
	fmt.Printf("brv: %v\n", brv)
	fmt.Println("brt2: ", brt2)
	fmt.Println("brk: ", brk)

	fmt.Println()
	fmt.Println("f: ", f)
	fmt.Printf("frt: %T\n", frt)
	fmt.Println("  frt: ", frt)
	fmt.Printf("frv: %v\n", frv)
	fmt.Println("frt2: ", frt2)
	fmt.Println("frk: ", frk)

	return n + 4
}

func main() {
	f(3)
}
