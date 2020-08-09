package graph

import (
  "testing"
)

func TestAddNode(t *testing.T) {
  var nodes = []Node { {"one", func() {

  }}, {"two", func() {

  }}}
  var narray = &nodes
  var graph = Graph{narray, nil}

  ans := AddNode(graph, Node{"three", nil})
  if !ans {
    t.Error("Failed to add node to graph")
  }
}