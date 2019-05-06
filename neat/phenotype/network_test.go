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

	if &net.nodes[0] == net.nodeMap[10]{
		t.Errorf("Node in net does not match the one in the map")
	}
}
