package models

import (
  "fmt"
  "math"
)

// Define a structure to hold the model's parameters
type Model struct {
  Weights []float64
  Bias float64
}

// Define a function to make predictions with the model
func (m *Model) Predict(inputs []float64) float64 {
  // Initialize the prediction value
  var prediction float64 = 0
  // Loop over the input values and weights and multiply them together
  for i := 0; i < len(inputs); i++ {
    prediction += inputs[i] * m.Weights[i]
  }
  // Add the bias to the prediction
  prediction += m.Bias
  // Return the prediction
  return prediction
}
