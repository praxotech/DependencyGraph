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

func AddNode(graph Graph, node Node) bool {
  var i int
  var nodes = graph.nodes
  for i = 1; i < len(*nodes); i++ {
    if (&(*nodes)[i] == &node) {
      return false
    }
  }

  var newNode = append((*nodes)[: len(*nodes)], node)
  graph.nodes = &newNode

  return true
}