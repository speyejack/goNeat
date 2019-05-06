package phenotype

import "testing"

func TestNetworkAddNodes(t *testing.T){
	net := &Network{}
	net.addNodes(10,12,13)


	if _,in := net.nodeMap[10]; !in {
		t.Errorf("Node failed to add with id: %v, map: %v", 10, net.nodeMap)
	}

	if _,in := net.nodeMap[12]; !in {
		t.Errorf("Node failed to add with id: %v, map: %v", 12, net.nodeMap)
	}

	if _,in := net.nodeMap[13]; !in {
		t.Errorf("Node failed to add with id: %v, map: %v", 13, net.nodeMap)
	}

	if &net.nodes[0] != net.nodeMap[10]{
		t.Errorf("Node in net does not match the one in the map\nExpected: %v\tGot: %v", &net.nodes[0], net.nodeMap[10])
	}
}

func TestNetworkAddAllNodes(t *testing.T){
	net := &Network{}
	ins := []int{1}
	hiddens := []int{3,6}
	outs := []int{10,5}

	net.addAllNodes(ins, outs, hiddens)
}

func TestNetworkSort(t *testing.T){
	ins := []int{1}
	hiddens := []int{7,3}
	outs := []int{10}

	links := []Link{
		{1,3,1.0},
		{3,7,1.0},
		{7,10,1.0}}

	net := BuildNetwork(ins, outs, hiddens, links)

	true_ids := []int{}
	for _,node := range net.nodes {
		true_ids = append(true_ids, node.id)
	}

	expected_ids := []int{1,3,7,10}
	failed := false
	for i,_ := range true_ids{
		if expected_ids[i] != true_ids[i] {
			failed = true
		}

	}
	if failed {
		t.Errorf("Got %v expected %v, sort failed", true_ids, expected_ids)
	}


}

func TestBuildNetwork(t *testing.T){
	ins := []int{1}
	hiddens := []int{94, 62, 30, 6}
	outs := []int{10}

	links := []Link{
		{1,30,1.0},
		{30,6,1.0},
		{30,94,1.0},
		{6,62,1.0},
		{94,62,1.0},
		{62,10,1.0}}

	net := BuildNetwork(ins, outs, hiddens, links)

	true_ids := []int{}
	for _,node := range net.nodes {
		true_ids = append(true_ids, node.id)
	}

	expected_ids := []int{1,30,94,6,62,10}
	failed := false
	for i,_ := range true_ids{
		if expected_ids[i] != true_ids[i] {
			failed = true
		}

	}
	if failed {
		t.Errorf("Got %v expected %v, sort failed", true_ids, expected_ids)
	}


}

func TestNetworkPropagation(t *testing.T){
	ins := []int{1}
	hiddens := []int{2}
	outs := []int{10}

	links := []Link{
		{1,2,1.0},
		{2,10,1.0}}

	net := BuildNetwork(ins, outs, hiddens, links)

	input := []float32{1.0}
	output := net.update(input)
	expected := float32(0.5)
	if output[0] <= expected {
		t.Errorf("Got %v expected greater than %v", output[0], expected)
	}
	old_output := output[0]

	input = []float32{0}
	output = net.update(input)
	expected = float32(0.5)
	if output[0] <= expected {
		t.Errorf("Got %v expected greater than %v", output[0], expected)
	}

	if output[0] >= old_output {
		t.Errorf("Got %v expected less than %v", output[0], old_output)
	}
}
