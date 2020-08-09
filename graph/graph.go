package graph

import (
  "bytes"
  "log"
)

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

var (
  buf bytes.Buffer
  logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func AddNode(graph *Graph, node Node) bool {
  var i int
  var nodes = (*graph).nodes
  for i = 0; i < len(*nodes); i++ {
    if NodesEqual((*nodes)[i], node) {
      logger.Printf("Node %+v already exists", node)
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

func PipesEqual(pipe1 Pipe, pipe2 Pipe) bool {
  return NodesEqual(pipe1.sink, pipe2.sink) && NodesEqual(pipe1.source, pipe2.source)
}