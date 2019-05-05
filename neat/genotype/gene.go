package genotype

type ConnectGene struct {
	in int
	out int
	weight float32
	enabled bool
	innov int
}

type NodeType int
const (
	Input NodeType = 0
	Output NodeType = 1
	Hidden NodeType = 2
)

type NodeGene struct {
	nodeType NodeType
}
