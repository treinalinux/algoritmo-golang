package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreveArea(f forma) {
	fmt.Printf(" A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	// return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreveArea(r)

	c := circulo{10}
	escreveArea(c)
}
