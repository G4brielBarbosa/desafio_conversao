package main

import "fmt"

const ptoEbulicaoK float64 = 373.0

func main() {

	ptoEbulicaoC := ptoEbulicaoK - 273

	fmt.Printf("O ponto de ebulição da agua em ºK é : %f\n", ptoEbulicaoK)
	fmt.Printf("O ponto de ebulição da agua em ºC é : %f\n", ptoEbulicaoC)
}
