package main

// import "math/rand"

type Layer struct {
	Neurons []*Neuron
}

func (layer *Layer) ConnectTo(to *Layer) {
	for _, n1 := range layer.Neurons {
		for _, n2 := range to.Neurons {
			n1.ConnectTo(n2)
		}
	}
}

func NewLayer(neurons int) *Layer {
	layer := &Layer{}
	for i := 0; i < neurons; i++ {
		layer.Neurons = append(layer.Neurons, &Neuron{Bias: 0})
	}

	return layer
}

func (layer *Layer) Set(input []float64) {
	for i, neuron := range layer.Neurons {
		neuron.State = input[i]
	}
}

func (layer *Layer) Reset() {
	for _, neuron := range layer.Neurons {
		neuron.State = 0
	}
}

func (layer *Layer) Process() {
	for _, neuron := range layer.Neurons {
		neuron.Process()
	}
}

func (layer *Layer) Output() []float64 {
	output := []float64{}
	for _, neuron := range layer.Neurons {
		output = append(output, neuron.State)
	}
	return output
}
