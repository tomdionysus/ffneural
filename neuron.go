package main

type Neuron struct {
	Forward  []*Link
	Backward []*Link
	Bias     float64
	State    float64
}

func (neuron *Neuron) ConnectTo(to *Neuron) {
	link := &Link{Weight: 1.0, From: neuron, To: to}
	neuron.Forward = append(neuron.Forward, link)
	to.Backward = append(neuron.Backward, link)
}

func (neuron *Neuron) Process() {
	for _, link := range neuron.Forward {
		link.Process()
	}
}
