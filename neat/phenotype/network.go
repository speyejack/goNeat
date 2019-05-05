package phenotype

import "github.com/speyejack/goNeat/neat/genotype/genome"

type Network struct {
	nodeMap map[int]*Node
	nodes []Node
	inputs []Node
	outputs []Node
}

