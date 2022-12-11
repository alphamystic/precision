package models

import (
  "os"
  "math/rand"

  "gonum.org/v1/gonum/mat"
)

type NeuralNet struct {
	inputs int
	hidden int
	outputs int
	weights *mat.Dense
	learningRate float64
	activation *Activation
}

type Activation struct {
	F, Fprime func(float64) float64
}

func New(inputs, hidden, outputs int, learningRate float64, activation *Activation) *NeuralNet {
	net := &NeuralNet {
		inputs: inputs,
		hidden: hidden,
		outputs: outputs,
		learningRate: learningRate,
		activation: activation,
	}
	net.weights = mat.NewDense(inputs+1, hidden+1, nil)
	rand.Seed(int64(os.Getpid()))
	net.weights.Apply(func(i, j int, v float64) float64 {
		return rand.Float64()
	}, net.weights)
	return net
}

func (nn *NeuralNet) FeedForward(inputs []float64) []float64 {
	if len(inputs) != nn.inputs {
		panic("invalid number of inputs")
	}
	inputs = append(inputs, 1.0)
	hiddenInputs := mat.NewVecDense(nn.inputs+1, inputs)
	hiddenOutputs := mat.NewVecDense(nn.hidden+1, nil)
	hiddenOutputs.MulVec(nn.weights, hiddenInputs)
	hiddenOutputs.Apply(func(i int, v float64) float64 {
		return nn.activation.F(v)
	}, hiddenOutputs)
	outputs := make([]float64, nn.outputs)
	for i := 0; i < nn.outputs; i++ {
		outputs[i] = nn.activation.F(hiddenOutputs.AtVec(i))
	}
	return outputs
}

func (nn *NeuralNet) Train(inputs, targets []float64) {
	outputs := nn.FeedForward(inputs)
	hiddenInputs := mat.NewVecDense(nn.inputs+1, inputs)
	hiddenOutputs := mat.NewVecDense(nn.hidden+1, nil)
	hiddenOutputs.MulVec(nn.weights, hiddenInputs)
	hiddenOutputs.Apply(func(i int, v float64) float64 {
		return nn.activation.F(v)
	}, hiddenOutputs)
	deltas := make([]float64, nn.outputs)
	for i := 0; i < nn.outputs; i++ {
		deltas[i] = nn.activation.Fprime(outputs[i]) * (targets[i] - outputs[i])
	}
	hiddenDeltas := make([]float64, nn.hidden)
	for i := 0; i < nn.hidden; i++ {
		error := 0.0
		for j := 0; j < nn.outputs; j++ {
			error += deltas[j] * nn.weights.At(i, j)
		}
		hiddenDeltas[i] = nn.activation.Fprime(hiddenOutputs.AtVec(i)) * error
	}
	for i := 0; i < nn.inputs+1; i++ {
		for j := 0; j < nn.hidden; j++ {
			change := hiddenDeltas[j] * hiddenInputs.AtVec(i)
			nn.weights.Set(i, j, nn.weights.At(i, j)+nn.learningRate*change)
		}
	}
	for i := 0; i < nn.hidden; i++ {
		for j := 0; j < nn.outputs; j++ {
			change := deltas[j] * hiddenOutputs.AtVec(i)
			nn.weights.Set(i, j, nn.weights.At(i, j)+nn.learningRate*change)
		}
	}
}
