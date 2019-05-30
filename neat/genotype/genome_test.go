package genotype

import "testing"

func TestGenomeNewGenomeNoHidden(t *testing.T){
	genomes := SpawnGenomes(1,1,[][]ProtoLink{{{0,1}}})
	genome := genomes[0]

	if len(genome.nodeGenes) != 2 {
		t.Errorf("Incorrect number of nodes created %v instead of %v\nNode Genes: %v", len(genome.nodeGenes), 2, genome.nodeGenes)
	}

	if len(genome.connectGenes) != 1 {
		t.Errorf("Incorrect number of nodes created %v instead of %v", len(genome.connectGenes), 1)
	}
}

func TestGenomeNewGenomeHidden(t *testing.T){
	genomes := SpawnGenomes(1,1,[][]ProtoLink{{{0,2},{2,1}}})
	genome := genomes[0]

	if len(genome.nodeGenes) != 3 {
		t.Errorf("Incorrect number of nodes created %v instead of %v\nNode Genes: %v", len(genome.nodeGenes), 3, genome.nodeGenes)
	}

	if len(genome.connectGenes) != 2 {
		t.Errorf("Incorrect number of nodes created %v instead of %v", len(genome.connectGenes), 2)
	}
}

func TestGenomeMutateConnectionWeight(t *testing.T){
	genomes := SpawnGenomes(1,1,[][]ProtoLink{{{0,1}}})
	genome := genomes[0]

	first_weight := genome.connectGenes[0].weight
	genome.mutateConnectionWeight(1.0)

	if genome.connectGenes[0].weight == first_weight {
		t.Errorf("Weight mutation didn't affect weight with value %v\n", first_weight)
	}

}

func TestGenomeMutateAddConnection(t *testing.T){
	genomes := SpawnGenomes(1,1,[][]ProtoLink{{}})
	genome := genomes[0]

	innov_chan := make(chan int,1)
	innov_chan<-10
	genome.mutateAddConnection(1.0,innov_chan)

	if len(genome.connectGenes) != 1 {
		t.Errorf("Incorrect number of nodes created %v instead of %v", len(genome.connectGenes), 1)
	}
}

func TestGenomeMutateAddNode(t *testing.T){
	genomes := SpawnGenomes(1,1,[][]ProtoLink{{{0,1}}})
	genome := genomes[0]

	innov_chan := make(chan int,2)
	innov_chan<-10
	innov_chan<-11
	genome.mutateAddNode(1.0,innov_chan)

	if len(genome.nodeGenes) != 3 {
		t.Errorf("Incorrect number of nodes created %v instead of %v\nNode Genes: %v", len(genome.nodeGenes), 3, genome.nodeGenes)
	}

	if len(genome.connectGenes) != 3 {
		t.Errorf("Incorrect number of links created %v instead of %v\nLink Genes: %v", len(genome.connectGenes), 3, genome.connectGenes)
	}
}
