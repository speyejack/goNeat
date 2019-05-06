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

	for i,_ := range new_nodes{
		new_nodes[i].id = ids[i]
		net.nodeMap[ids[i]] = &new_nodes[i]
	}
}

func (net *Network) addAllNodes(inIDs []int, outIDs []int, hiddenIDs []int){
	net.nodeMap = make(map[int]*Node)
	net.nodes = make([]Node, 0, len(inIDs) + len(outIDs) + len(hiddenIDs))

	net.addNodes(inIDs...)
	net.addNodes(hiddenIDs...)
	net.addNodes(outIDs...)
	net.inputs = net.nodes[:len(inIDs)]
	net.hidden = net.nodes[len(inIDs):len(inIDs) + len(hiddenIDs)]
	net.outputs = net.nodes[len(inIDs) + len(hiddenIDs):]
}

func (net *Network) addLink(link Link){
	in, out := link.in, link.out
	nodeLink := NodeLink{net.nodeMap[in], net.nodeMap[out], link.weight}

	net.nodeMap[in].outLinks = append(net.nodeMap[in].outLinks, nodeLink)
	net.nodeMap[out].inLinks = append(net.nodeMap[out].inLinks, nodeLink)
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
				inRank += old_ranks[link.in.id]
			}
			inRank /= float32(len(node.inLinks))

			outRank := float32(0.0)
			for _,link := range node.outLinks {
				outRank += old_ranks[link.out.id]
			}
			outRank /= float32(len(node.outLinks))

			new_ranks[id] = (inRank + outRank)/2
		}
		old_ranks, new_ranks = new_ranks, old_ranks
	}
	sorter := nodeSorter{net.nodes, old_ranks}
	sort.Sort(&sorter)

}

func (net *Network) build(inIDs []int, outIDs []int, hiddenIDs []int, links []Link) {

	net.addAllNodes(inIDs, outIDs, hiddenIDs)

	for _,link := range links{
		net.addLink(link)
	}

	net.sortNodes()
}

func (net *Network) setInputs(inputs []float32){
	for i,value := range inputs {
		net.inputs[i].value = value;
	}
}

func (net *Network) getOutputs() (outputs []float32){
	outputs = make([]float32, 0, len(net.outputs))
	for _,node := range net.outputs {
		outputs = append(outputs,node.value)
	}
	return
}

func (net *Network) propagate(){
	nodes := net.nodes[len(net.inputs):]
	for i,_ := range nodes{
		nodes[i].update()
	}
}

func (net *Network) update(inputs []float32) []float32{
	net.setInputs(inputs)
	net.propagate()
	return net.getOutputs()
}

func BuildNetwork(inIDs []int, outIDs []int, hiddenIDs []int, links []Link) (net *Network){
	net = &Network{}
	net.build(inIDs, outIDs, hiddenIDs, links)
	return
}
