package models



// Define a function to calculate the mean squared error
func MeanSquaredError(predictions []float64, targets []float64) float64 {
  // Initialize the mean squared error
  var mse float64 = 0
  // Loop over the predictions and targets and calculate the mean squared error
  for i := 0; i < len(predictions); i++ {
    mse += math.Pow(predictions[i] - targets[i], 2)
  }
  // Divide the mean squared error by the number of predictions to get the mean
  mse /= float64(len(predictions))
  // Return the mean squared error
  return mse
}

// Define a function to apply the sigmoid function to a value
func Sigmoid(x float64) float64 {
  return 1 / (1 + math.Exp(-x))
}


func SigmoidPrime(x float64) float64 {
	return Sigmoid(x) * (1.0 - Sigmoid(x))
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func TanhPrime(x float64) float64 {
	return 1.0 - math.Pow(Tanh(x), 2)
}

func ReLU(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func ReLUPrime(x float64) float64 {
	if x > 0 {
		return 1.0
	}
	return 0.0
}
