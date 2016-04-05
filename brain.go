package deanio

import (
	"fmt"
	"github.com/fogleman/nes/nes"
)

var seconds int
var ticks uint64
var currentGen *Genome
var currentGenI int
var currentPool *Pool

func initBrain() {
	currentPool = generatePool()
	currentGen = currentPool.genomes[currentGenI+1]
}

func SaveInit(console *nes.Console) {
	console.SaveState("womp/" + currentPool.baseFilename)
}

func Step(console *nes.Console) {
	if ticks == TOTAL_TICKS-1 {
		console.SaveState("womp/" + currentGen.filename)
		ticks = 0
		if currentGenI >= 10 {
			currentGen = currentPool.getBest()
			fmt.Println("BEST")
			console.LoadState("womp/" + currentGen.filename)
		}
	}
	if ticks == 0 {
		console.Reset()
		console.LoadState(currentPool.baseFilename)
		currentGen = currentPool.genomes[currentGenI]
		currentGenI++
		fmt.Println(currentGenI)
		fmt.Println(currentGen.filename)
		fmt.Println(currentPool.baseFilename)
	}
	currentGen.outputs[ticks].memorySum = ReadMem(console.RAM)
	ticks += 1
}

func ReadController() (controllerOut [8]bool) {
	controllerOut = currentGen.outputs[ticks].controllerInput
	printOut := ""
	for i := 0; i <= 7; i++ {
		if controllerOut[i] {
			switch i {
			case 0:
				printOut += "A "
			case 1:
				printOut += "B "
			case 2:
				printOut += "Select "
			case 3:
				printOut += "Start "
			case 4:
				printOut += "Up "
			case 5:
				printOut += "Down "
			case 6:
				printOut += "Left "
			case 7:
				printOut += "Right "
			}
		}
	}

	updateController(printOut)

	return
}
