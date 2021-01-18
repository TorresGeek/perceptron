package main

import (
	"fmt"
	perceptronPackage "perceptron/perceptron"
)

func main() {
	orPerceptron := perceptronPackage.MakeBasePerceptron()
	orPerceptron.TrainORGate()

	andPerceptron := perceptronPackage.MakeBasePerceptron()
	andPerceptron.TrainANDGate()

	fmt.Printf("[OrGate]\n")
	fmt.Printf("0 | 0 :  %d\n", orPerceptron.CalculateResult(0, 0))
	fmt.Printf("0 | 1 :  %d\n", orPerceptron.CalculateResult(0, 1))
	fmt.Printf("1 | 0 :  %d\n", orPerceptron.CalculateResult(1, 0))
	fmt.Printf("1 | 1 :  %d\n", orPerceptron.CalculateResult(1, 1))

	fmt.Printf("\n\n[AndGate]\n")
	fmt.Printf("0 | 0 :  %d\n", andPerceptron.CalculateResult(0, 0))
	fmt.Printf("0 | 1 :  %d\n", andPerceptron.CalculateResult(0, 1))
	fmt.Printf("1 | 0 :  %d\n", andPerceptron.CalculateResult(1, 0))
	fmt.Printf("1 | 1 :  %d\n", andPerceptron.CalculateResult(1, 1))
}
