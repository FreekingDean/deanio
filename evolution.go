package deanio

import (
	"strconv"
	"time"
)

const (
	NUM_GENOMES   = 25
	MUTATION_RATE = 0.02
	TOTAL_TICKS   = 6000
)

type Pool struct {
	genomes      []*Genome
	baseGenome   *Genome
	baseFilename string
}

type Genome struct {
	outputs   map[uint64]*Stat
	filename  string
	memorySum uint64
}

type Stat struct {
	memorySum       uint64
	controllerInput [8]bool
}

func generatePool() *Pool {
	pool := &Pool{
		baseGenome:   generateGenome(nil),
		baseFilename: "pool_base.dat",
	}

	pool.genomes = make([]*Genome, 0)
	for i := 0; i < NUM_GENOMES; i++ {
		pool.genomes = append(pool.genomes, generateGenome(pool.baseGenome))
	}

	return pool
}

func generateGenome(base *Genome) *Genome {
	genome := &Genome{}
	genome.outputs = make(map[uint64]*Stat)
	genome.filename = strconv.FormatInt(time.Now().UnixNano(), 16) + ".state"

	for i := 0; i < TOTAL_TICKS; i++ {
		var controllerInput [8]bool
		if calcProb() < MUTATION_RATE || base == nil {
			if base == nil {
				var empty [8]bool
				controllerInput = generateControllers(empty, true)
			} else {
				controllerInput = generateControllers(base.outputs[uint64(i)].controllerInput, false)
			}
		} else {
			controllerInput = base.outputs[uint64(i)].controllerInput
		}

		genome.outputs[uint64(i)] = &Stat{
			memorySum:       0,
			controllerInput: controllerInput,
		}
	}

	return genome
}

func (g *Genome) totalSum() {
	g.memorySum = 0
	for _, stat := range g.outputs {
		g.memorySum += stat.memorySum
	}
}

func (p *Pool) getBest() *Genome {
	best := p.genomes[0]
	for _, genome := range p.genomes {
		if best.memorySum < genome.memorySum {
			best = genome
		}
	}
	return best
}
