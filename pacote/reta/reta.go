package main

import "math"

// Pacotes são a forma de modularizar o código
// podemos criar libs para reuso

// Iniciando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)
// Iniciando com letra minúscula é PRIVADO (visível somente dentro do pacote)

//Exemplo
// Ponto - Público
// ponto ou _Ponto - Privado

// precisamos colocar um comentário acima da estrutura pública para funcionar
// Ex: Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return cx, cy
}

// Distância é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
