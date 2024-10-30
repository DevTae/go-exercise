package main

import (
	"fmt"
	"math"
)


type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat 은 Abser 인터페이스를 구현했으니 가능함
	a = &v // *Vertex 는 Abser 인터페이스를 구현했으니 가능함

	a = v  // v 는 Vertex 타입임 (*Vertex 타입이 아님), value receiver 는 구현안했으니 에러!!

	fmt.Println(a.Abs())
}
