package main

type Network struct {
	Layers []*Layer
}

func NewNetwork(inputs int, hidden []int, outputs int) *Network {
	network := &Network{}
	network.Layers = append(network.Layers, NewLayer(inputs))
	for i := 0; i < len(hidden); i++ {
		network.Layers = append(network.Layers, NewLayer(hidden[i]))
	}
	network.Layers = append(network.Layers, NewLayer(outputs))

	for i := 0; i < len(network.Layers)-1; i++ {
		network.Layers[i].ConnectTo(network.Layers[i+1])
	}

	return network
}

func (network *Network) Input(inputs []float64) {
	network.Layers[0].Set(inputs)
}

func (network *Network) Output() []float64 {
	return network.Layers[len(network.Layers)].Output()
}

func (network *Network) Reset() {
	for _, layer := range network.Layers {
		layer.Reset()
	}
}

func (network *Network) Process() {
	for _, layer := range network.Layers {
		layer.Process()
	}
}
