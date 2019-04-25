package main

import "unsafe"
import "fmt"
import "reflect"

func P(t string, n uintptr) {
	fmt.Printf("%s :  %d\n", t, n)
}
func P2(t reflect.Type, v reflect.Value) {
	fmt.Println("reflect t :", t, ", v :", v)
}

func f(x float32) float32 {
	return x + 1
}

func main() {
	fmt.Println("Sizes")

	{
		x := 6
		t := reflect.TypeOf(x)
		v := reflect.ValueOf(x)
		P2(t, v)
	}
	{
		x := func(i int, x float32) int { return i + int(x) }
		t := reflect.TypeOf(x)
		v := reflect.ValueOf(x)
		P2(t, v)
	}
	{
	}
	{
	}
	{
	}
	{
		var x int
		P("int", unsafe.Sizeof(x))
	}
	{
		x := 1
		P("x:=1", unsafe.Sizeof(x))
	}
	{
		x := 1.0
		P("x:=1.0", unsafe.Sizeof(x))
	}
	{
		x := 1.0 + 2.0i
		P("x:=1.0+2.0i", unsafe.Sizeof(x))
	}
	{
		var x [6]int
		P("[6]int", unsafe.Sizeof(x))
	}
	{
		var x []int
		P("[]int", unsafe.Sizeof(x))
	}
	{
		var x float32
		P("float32", unsafe.Sizeof(x))
	}
	{
		type S = float32
		var x S
		P("float32", unsafe.Sizeof(x))
	}
	{
		type S float32
		var x S
		P("Nfloat32", unsafe.Sizeof(x))
	}
	{
		var x []float32
		P("[]float32", unsafe.Sizeof(x))
	}
	{
		var x [3]float32
		P("[3]float32", unsafe.Sizeof(x))
	}
	{
		var x [7][3]float32
		P("[7][3]float32", unsafe.Sizeof(x))
		x[3][1] = 1.3
		fmt.Println("    ", x)
	}
	{
		var x [3][7]float32
		P("[3][7]float32", unsafe.Sizeof(x))
	}
	{
		var x []float64
		P("[]float64", unsafe.Sizeof(x))
	}
	{
		var x int16
		P("int16", unsafe.Sizeof(x))
	}
	{
		var x int
		P("int", unsafe.Sizeof(x))
	}
	{
		type S = struct{}
		var x S
		P("struct{}", unsafe.Sizeof(x))
	}
	{
		type S = interface{}
		var x S
		P("interface{}", unsafe.Sizeof(x))
	}
	{
		type S = struct{ int }
		var x S
		P("struct{int}", unsafe.Sizeof(x))
	}
	{
		type S = struct {
			a int
			b byte
			c int
			d byte
		}
		var x S
		P("struct{int;byte;int;byte}", unsafe.Sizeof(x))
	}
	{
		type S = struct {
			a int
			c int
			b byte
			d byte
		}
		var x S
		P("struct{int;int;byte;byte}", unsafe.Sizeof(x))
	}
	{
		type S = struct {
			a int
			c int
			b byte
			d byte
		}
		var x S
		P("struct{int;int;byte;byte}", unsafe.Sizeof(x))
		x.c = 2
		fmt.Println("    ", x)
	}
	{
		var i int
		x := &i
		P("*int", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		type S = func(int) int
		var x S
		P("func(int)int", unsafe.Sizeof(x))
	}
	{
		type S = func(int, int) int
		var x S
		P("func(int,int)int", unsafe.Sizeof(x))
	}
	{
		x := f
		P("f", unsafe.Sizeof(x))
		fmt.Println("    ", f)
	}
	{
		x := main
		P("main", unsafe.Sizeof(x))
		fmt.Println("    ", main)
	}
	{
		type S = map[string]int
		var x S
		P("map[string]int", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		x := make(map[string]int)
		x["aa"] = 3
		P("map[string]int", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		x := "string"
		P("string", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		type S = string
		var x S = "abc"
		P("string", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		type S = chan int
		var x S
		P("chan int", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		type S = func(float32) float32
		var x S
		P("func(float32)float32", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		var x complex64 = complex(2.1, 3.7)
		P("complex64", unsafe.Sizeof(x))
		fmt.Println("    ", x)
	}
	{
		var x complex128
		P("complex128", unsafe.Sizeof(x))
	}
}
