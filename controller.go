package deanio

var lastPressed [8]bool

func generateControllers(controllerIn [8]bool, new bool) (controllerOut [8]bool) {
	for i := 0; i <= 7; i++ {
		if new || calcProb() < MUTATION_RATE {
			controllerOut[i] = shouldPress(lastPressed[i])
		} else {
			controllerOut[i] = controllerIn[i]
		}
	}

	if controllerOut[6] && controllerOut[7] {
		controllerOut[6] = false
		controllerOut[7] = false
	}

	if controllerOut[4] && controllerOut[5] {
		controllerOut[4] = false
		controllerOut[5] = false
	}
	lastPressed = controllerOut
	return
}

func shouldPress(isPressed bool) bool {
	prob := calcProb()
	return (isPressed && prob < 0.9) || prob < 0.1
}
