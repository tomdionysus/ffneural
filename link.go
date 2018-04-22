package main

import "math"

type Link struct {
	Weight float64
	To     *Neuron
	From   *Neuron
}

func sigmoid(x float64) float64 {
	return (1.0 / (1.0 + math.Exp(-x)))
}

func ReLU(x float64) float64 {
	return math.Max(0, x)
}

func (link *Link) Process() {
	link.To.State += link.From.State*sigmoid(link.Weight) + link.From.Bias
}
