package formas

import (
	"math"
)

// Forma uma interface que recebe uma area do tipo float64
type Forma interface {
	area() float64
}

// Retangulo uma struct que recebe al e largura do tipo float64
type Retangulo struct {
	altura  float64
	largura float64
}

// Area do Retangulo
func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

// Circulo uma struct que recebe o raio
type Circulo struct {
	raio float64
}

// Area do  Circulo
func (c Circulo) Area() float64 {
	// return math.Pi * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}
