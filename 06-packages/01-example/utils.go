package main

import "math"

// Iniciando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)
// Iniciando com letra minúscula ou _ é PRIVADO (visível apenas dentro do mesmo pacote)

// Isso se aplica para
// 1. Funções
// 2. Nome de struct
// 3. Propriedades do struct
// 4. Tipos
// 5. Constantes e Variáveis

type Coord struct {
	x float64
	y float64
}

func legs(a, b Coord) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distance(a, b Coord) float64 {
	cx, cy := legs(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
