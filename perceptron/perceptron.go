package perceptron

/*Perceptron : PerceptronData*/
type Perceptron struct {
	Lr    float32
	Bias  float32
	Pesos []float32
}

/*MakeBasePerceptron :MakeBasePerceptron*/
func MakeBasePerceptron() Perceptron {
	perceptronInstance := new(Perceptron)
	perceptronInstance.Bias = 1
	perceptronInstance.Lr = 1
	perceptronInstance.Pesos = make([]float32, 3)
	perceptronInstance.Pesos[0] = 0.5
	perceptronInstance.Pesos[1] = 0.5
	perceptronInstance.Pesos[2] = 0.5

	return *perceptronInstance
}

/*CalculateResult : CalculateResult*/
func (Per Perceptron) CalculateResult(x float32, y float32) int32 {
	value := ((x * Per.Pesos[0]) + (y * Per.Pesos[1]) + (Per.Bias * Per.Pesos[2]))
	if value > 0 {
		return 1
	}

	return 0
}

/*PerceptronProcess : PerceptronProcess */
func (Per Perceptron) PerceptronProcess(x1 float32, x2 float32, respuestaReal float32) {
	PosibleRespuesta := ((x1 * Per.Pesos[0]) + (x2 * Per.Pesos[1]) + (Per.Bias * Per.Pesos[2]))
	if PosibleRespuesta > 0 {
		PosibleRespuesta = 1
	} else {
		PosibleRespuesta = 0
	}

	errorCalculado := respuestaReal - PosibleRespuesta

	Per.Pesos[0] += (errorCalculado * x1 * Per.Lr)
	Per.Pesos[1] += (errorCalculado * x2 * Per.Lr)
	Per.Pesos[2] += (errorCalculado * Per.Bias * Per.Lr)
}

/*TrainORGate : TrainORGate*/
func (Per Perceptron) TrainORGate() {
	for i := 0; i < 10000; i++ {
		Per.PerceptronProcess(1, 1, 1)
		Per.PerceptronProcess(1, 0, 1)
		Per.PerceptronProcess(0, 1, 1)
		Per.PerceptronProcess(0, 0, 0)
	}
}

/*TrainANDGate : TrainANDGate*/
func (Per Perceptron) TrainANDGate() {
	for i := 0; i < 10000; i++ {
		Per.PerceptronProcess(1, 1, 1)
		Per.PerceptronProcess(1, 0, 0)
		Per.PerceptronProcess(0, 1, 0)
		Per.PerceptronProcess(0, 0, 0)
	}
}
