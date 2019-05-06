package phenotype

import "math"

type Node struct {
	id int
	inLinks []NodeLink
	outLinks []NodeLink
	value float32
}

func (node *Node) update (){
	value := float32(0)
	for _,link := range node.inLinks {
		value += link.in.value * link.weight
	}
	node.value = sigmoid(value)
}

func sigmoid(x float32) float32{
	return float32(1/(1+ math.Exp(-float64(x))))
}
