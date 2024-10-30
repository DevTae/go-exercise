package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Vertex 에서의 값 반환 함수
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Vertex 에서의 값 수정 함수 (Receiver Function Test Struct)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{X:3, Y:4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
