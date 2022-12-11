package main

import (
  "fmt"
  "github.com/alphamystic/precision/ml/models"
)
func main() {
  // Define some input data and corresponding targets
  var inputs = [][]float64{
    {0, 0},
    {0, 1},
    {1, 0},
    {1, 1},
  }
  var targets = []float64{0, 1, 1, 0}

  // Create a new model with random weights and biases
  var model = models.Model{
    Weights: [][]float64{
      {0.5, -0.5},
      {0.5, 0.5},
    },
    Biases: []float64{0, 0},
  }

  // Make predictions with the model
  var predictions []float64
  for i := 0; i < len(inputs); i++ {
    predictions = append(predictions, model.Predict(inputs[i])...)
  }

  // Print the predictions and the mean squared error
  fmt.Println("Predictions:", predictions)
  fmt.Println("Mean Squared Error:", models.MeanSquaredError(predictions, targets))
}
