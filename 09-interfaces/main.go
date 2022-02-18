package main

import (
	"fmt"
)

//interface == contrato, quem implementar deverá cumprir contrato
type transporter interface {
	calculateFreight(value float32) float32
}

type correios struct {}

func (c correios) calculateFreight(value float32) float32 {
	return value * 0.3;
}

type fedex struct {}

//f fedex virou transporter assim que implementou função calculateFreight
func (f fedex) calculateFreight(value float32) float32 {
	return value * 0.5;
}

type novaTransportadora struct {}

func (n novaTransportadora) calculateFreight(value float32) float32 {
	return value * 0.8
}

func main() {
	value := float32(53.1)

	freights := make(map[string]float32)

	freights["Correios"] = getFreights(correios{}, value)
	freights["Fedex"] = getFreights(fedex{}, value)

	fmt.Println(freights)

}

func getFreights(t transporter, value float32) float32 {
	return t.calculateFreight(value)
}