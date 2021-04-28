package service

import "fmt"

// Node is a module with its version, and which modules use it (upstreams) and which module it depends on (downstreams)
type Node struct {
	VersionedModule VersionedModule
	Downstreams     []*Node // all the dependencies from this node
	Upstreams       []*Node // all the other nodes that depend on this node
}

func (n *Node) Print() {
	fmt.Print(n.VersionedModule.ToString(), " ")
	fmt.Print(" [upstreams]: ")
	for _, upNode := range n.Upstreams {
		fmt.Print(upNode.VersionedModule.ToString() + " ")
	}
	fmt.Print("[downstreams]: ")
	for _, downNode := range n.Downstreams {
		fmt.Print(downNode.VersionedModule.ToString() + " ")
	}
}

// AnalyzedTree contains all nodes with their upstream and downstream parsed.
// From this tree one can search for the DependencyPath from root to a given module@version
type AnalyzedTree struct {
	Root *Node
	SearchMap map[string]map[string]*Node // key1 = module, key2 = version
}

func (t *AnalyzedTree) Print() {
	fmt.Println("root: ")
	t.Root.Print()
	fmt.Println()

	fmt.Println("searchMap: ")
	for m, ns := range t.SearchMap {
		fmt.Print(m + "->")
		for v, n := range ns {
			fmt.Print(v + "->")
			n.Print()
		}
		fmt.Println()
	}
}

func Analyze(inputs []Input) *AnalyzedTree {
	if len(inputs) == 0 {
		return nil
	}

	tree := AnalyzedTree{
		Root: nil,
		SearchMap: map[string]map[string]*Node{},
	}

	for i, input := range inputs {
		// first row is assumed to be the root
		if i == 0 {
			root := Node {
				VersionedModule: input.Dependant,
			}
			tree.Root = &root
		}

		// make sure dependant node and dependency node appear in search map
		m := input.Dependant.Module
		v := input.Dependant.Version
		if _, found := tree.SearchMap[m]; !found {
			tree.SearchMap[m] = map[string]*Node{}
		}
		if _, found := tree.SearchMap[m][v]; !found {
			tree.SearchMap[m][v] = &Node{ VersionedModule: input.Dependant }
		}
		dependantNode := tree.SearchMap[m][v]

		m = input.Dependency.Module
		v = input.Dependency.Version
		if _, found := tree.SearchMap[m]; !found {
			tree.SearchMap[m] = map[string]*Node{}
		}
		if _, found := tree.SearchMap[m][v]; !found {
			tree.SearchMap[m][v] = &Node{ VersionedModule: input.Dependency }
		}
		dependencyNode := tree.SearchMap[m][v]

		// process wiring
		// dependant => (downstream) => dependency
		downstreamExists := false
		for i := range dependantNode.Downstreams {
			if dependantNode.Downstreams[i] == dependencyNode {
				downstreamExists = true
				break
			}
		}
		if !downstreamExists {
			dependantNode.Downstreams = append(dependantNode.Downstreams, dependencyNode)
		}

		// dependency => (upstream) => dependant
		upstreamExists := false
		for i := range dependencyNode.Upstreams {
			if dependencyNode.Upstreams[i] == dependantNode {
				upstreamExists = true
				break
			}
		}
		if !upstreamExists {
			dependencyNode.Upstreams = append(dependencyNode.Upstreams, dependantNode)
		}

		if i == 0 {
			// treat first input to contain the root
			tree.Root = dependantNode
		}
	}

	return &tree
}
