package graph

import (
  "testing"
)

func TestAddNode(t *testing.T) {
  var nodes = []Node { {"one", func() {}}, {"two", func() {}}}
  t.Log(nodes)
  var narray = &nodes
  var graph = Graph{narray, nil}
  t.Log(*graph.nodes)

  ans := AddNode(&graph, Node{"three", func() {}})
  if !ans {
    t.Error("Failed to add node to graph")
  }
  t.Log(*graph.nodes)

  ans = AddNode(&graph, nodes[0])
  if ans {
    t.Error("Error: Duplicate node added to graph")
  }
  t.Log(*graph.nodes)
}