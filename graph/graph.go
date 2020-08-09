package graph

type Node struct {
  name string
  operation func()
}

type Pipe struct {
  source Node
  sink Node
}

type Graph struct {
  nodes *[]Node
  pipes *[]Pipe
}

func AddNode(graph *Graph, node Node) bool {
  var i int
  var nodes = (*graph).nodes
  for i = 0; i < len(*nodes); i++ {
    if NodesEqual((*nodes)[i], node) {
      return false
    }
  }

  var newNodes = append((*nodes)[: len(*nodes)], node)
  (*graph).nodes = &newNodes

  return true
}

func NodesEqual(node1 Node, node2 Node) bool {
  return node1.name == node2.name
}