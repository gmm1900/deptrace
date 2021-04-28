package service

import (
	"fmt"
	"log"
)

// DependencyPath is a list that starts from root and ends with a particular module@version searched
// each element is a string of the form module@version
type DependencyPath []string

func (p *DependencyPath) Print() {
	for i, identifier := range *p {
		// padding
		for j := 0; j <= i * 2; j++ {
			fmt.Print(" ")
		}
		fmt.Println(identifier)
	}
}

func (t *AnalyzedTree) SearchPath(versionedModuleStr string) []DependencyPath {
	versionedModule, err := parseVersionedModule(versionedModuleStr)
	if err != nil {
		log.Fatalf("invalid versioned module: %v", versionedModuleStr)
	}
	nodesMap, found := t.SearchMap[versionedModule.Module]
	if !found {
		return nil
	}
	var nodeList []*Node
	if versionedModule.Version != "" {
		node, found := nodesMap[versionedModule.Version]
		if !found {
			return nil
		}
		nodeList = append(nodeList, node)
	} else {
		for _, node := range nodesMap {
			nodeList = append(nodeList, node)
		}
	}

	var result []DependencyPath
	for _, node := range nodeList {
		var perNodeResult []DependencyPath
		t.searchNode(node, DependencyPath{}, &perNodeResult)
		result = append(result, perNodeResult...)
	}
	return result
}

// searchNode recursively trace all the paths from a node, up to root or up to when there is no more upstream
func (t *AnalyzedTree) searchNode(node *Node, currentPath DependencyPath, allPaths *[]DependencyPath) {
	// does this path have circular dependency? if so stop this path.
	nodeStr := node.VersionedModule.ToString()
	for _, n := range currentPath {
		if n == nodeStr {
			// encountered circular dependency: this node has already been in the path; retract
			return
		}
	}
	currentPath = append([]string{nodeStr}, currentPath...)
	if len(node.Upstreams) == 0 {
		// no more upstream
		if node == t.Root {
			*allPaths = append(*allPaths, currentPath)
		} else {
			// add a warning node that the path didn't trace up to the root
			currentPath = append([]string{"[NOT ROOT]"}, currentPath...)
			*allPaths = append(*allPaths, currentPath)
		}
		return
	}

	for _, upstreamNode := range node.Upstreams {
		// recursively go upstream, for each upstream node
		t.searchNode(upstreamNode, currentPath, allPaths)
	}
}