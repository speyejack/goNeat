package genotype

import "math/rand"

type Genome struct {
	nodeGenes []NodeGene
	connectGenes []ConnectGene
}

func max(a int, b int) int{
	if a >= b{
		return a
	} else {
		return b
	}
}

func SpawnGenomes(inputs int, outputs int, links [][]ProtoLink) (genomes []Genome){
	innov_map := make(map[ProtoLink]int)
	done_chan := make(chan int)
	innov_chan := make(chan int)

	go func() {
		x := 0
		for {
			select {
			case innov_chan<- x:
				x++
			case <-done_chan:
				return
			}
		}
	}()
	defer func(){done_chan<-1}()

	for _,protolink := range links {
		genome := NewGenome(1, 1, protolink, innov_map, innov_chan)
		genomes = append(genomes, genome)
	}

	return
}

func NewGenome(inputs int, outputs int, links []ProtoLink,
	innovations map[ProtoLink]int, innovator chan int) (genome Genome){

	genome = Genome{nil, nil}
	for i := 0; i < inputs; i++{
		genome.nodeGenes = append(genome.nodeGenes, NodeGene{Input})
	}
	for i := 0; i < outputs; i++{
		genome.nodeGenes = append(genome.nodeGenes, NodeGene{Output})
	}

	for _,link := range links {
		for i := 0; i < max(link.in,link.out) - len(genome.connectGenes) - 1; i++{
			genome.nodeGenes = append(genome.nodeGenes, NodeGene{Hidden})
		}
		innov, found := innovations[link]
		if !found {
			innovations[link] = <-innovator
			innov = innovations[link]
		}
		genome.connectGenes = append(genome.connectGenes, ConnectGene{link.in, link.out, 1.0, true, innov})
	}
	return
}


func (genome *Genome) mutateConnectionWeight(prob float32){
	for i,_ := range genome.connectGenes {
		chance_val := rand.Float32()
		if chance_val < prob {
			genome.connectGenes[i].weight += float32(rand.NormFloat64())
		}
	}
}

func (genome *Genome) mutateAddConnection(prob float32, innovator chan int){
	chance_val := rand.Float32()
	if chance_val < prob {
		links := make(map[ProtoLink]bool)
		for _,link := range genome.connectGenes {
			links[ProtoLink{link.in,link.out}] = true
			links[ProtoLink{link.out,link.in}] = true
		}

		unconnected := make([]ProtoLink, 0, 5)
		for i:= 0; i < len(genome.nodeGenes); i++ {
			for j:= i+1; j < len(genome.nodeGenes); j++ {
				if _,linked := links[ProtoLink{i,j}];!linked {
					unconnected = append(unconnected, ProtoLink{i,j})
				}

			}
		}
		if len(unconnected) == 0 {
			return
		}
		node_index := rand.Intn(len(unconnected))
		innovation := <-innovator
		new_protolink := unconnected[node_index]
		gene := ConnectGene{new_protolink.in, new_protolink.out, 1.0, true, innovation}
		genome.connectGenes = append(genome.connectGenes,gene)
	}
}

func (genome *Genome) mutateAddNode(prob float32, innovator chan int){
	active := make([]*ConnectGene,0, len(genome.connectGenes))
	for _,link := range genome.connectGenes {
		if link.enabled {
			active = append(active, &link)
		}
	}

	link_index := rand.Intn(len(active))
	link := &genome.connectGenes[link_index]
	link.enabled = false

	genome.nodeGenes = append(genome.nodeGenes, NodeGene{Hidden})
	new_gene_num := len(genome.nodeGenes)-1

	innovation := <-innovator
	gene := ConnectGene{link.in, new_gene_num, 1.0, true, innovation}
	genome.connectGenes = append(genome.connectGenes, gene)

	innovation = <-innovator
	gene = ConnectGene{new_gene_num, link.out, link.weight, true, innovation}
	genome.connectGenes = append(genome.connectGenes, gene)
	}
