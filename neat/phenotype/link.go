package phenotype

type Link struct {
	in int
	out int
	weight float32
}

type NodeLink struct {
	in *Node
	out *Node
	weight float32
}
