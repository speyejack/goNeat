package phenotype

import (
	"sort"
)

type Network struct {
	nodeMap map[int]*Node
	nodes []Node
	inputs []Node
	outputs []Node
	hidden []Node
}

type nodeSorter struct {
	nodes []Node
	ranks map[int]float32
}

func (sorter *nodeSorter) Len() int { return len(sorter.nodes) }

func (sorter *nodeSorter) Less(i,j int) bool { return sorter.ranks[sorter.nodes[i].id] < sorter.ranks[sorter.nodes[j].id] }

func (sorter *nodeSorter) Swap(i,j int) { sorter.nodes[i], sorter.nodes[j] = sorter.nodes[j], sorter.nodes[i] }

func (net *Network) addNodes(ids ...int){
	addition_nodes := make([]Node, len(ids))

	pre_length := len(net.nodes)

	net.nodes = append(net.nodes, addition_nodes...)
	new_nodes := net.nodes[pre_length:]

	if net.nodeMap == nil {
		net.nodeMap = make(map[int]*Node)
	}

	for i,node := range new_nodes{
		node.id = ids[i]
		net.nodeMap[node.id] = &node;
	}
}

func (net *Network) addAllNodes(inIDs []int, outIDs []int, hiddenIDs []int){
	net.addNodes(inIDs...)
	net.inputs = net.nodes[:]
	net.addNodes(outIDs...)
	net.outputs = net.nodes[len(net.inputs):]
	net.addNodes(hiddenIDs...)
	net.hidden = net.nodes[len(net.inputs) + len(net.outputs):]
}

func (net *Network) addLink(in int, out int, weight float32){
	link := Link{in, out, weight}

	net.nodeMap[in].inLinks = append(net.nodeMap[in].inLinks, link)
	net.nodeMap[out].outLinks = append(net.nodeMap[out].outLinks, link)
}

func (net *Network) sortNodes(){
	new_ranks := make(map[int]float32)
	old_ranks := make(map[int]float32)

	for _, node := range net.inputs {
		new_ranks[node.id] = 0
		old_ranks[node.id] = 0
	}

	for _, node := range net.outputs {
		new_ranks[node.id] = 1
		old_ranks[node.id] = 1
	}

	for _, node := range net.hidden {
		new_ranks[node.id] = 0.5
		old_ranks[node.id] = 0.5
	}

	for i := 0; i < 10; i++{
		for _, node := range net.hidden{
			id := node.id
			inRank := float32(0.0)
			for _,link := range node.inLinks {
				inRank += old_ranks[link.in]
			}
			inRank /= float32(len(node.inLinks))

			outRank := float32(0.0)
			for _,link := range node.outLinks {
				outRank += old_ranks[link.out]
			}
			outRank /= float32(len(node.outLinks))

			new_ranks[id] = (inRank + outRank)/2
		}
		old_ranks, new_ranks = new_ranks, old_ranks
	}
	sorter := nodeSorter{net.nodes, old_ranks}
	sort.Sort(&sorter)
}

