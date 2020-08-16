package graph

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	graph := NewGraph()
	node := NewNode("one")
	node.SetOperation(func() {})
	graph.AddNode(*node)
	node = NewNode("two")
	node.SetOperation(func() {})
	graph.AddNode(*node)
	t.Log((*graph).nodes)

	ans := graph.AddNode(*NewNode("two"))
	if ans {
		t.Error("Error: duplicate node is added")
	}

	ans = graph.AddNode(*NewNode("three"))
	if !ans {
		t.Error("Error: failed adding new node")
	}
	t.Log((*graph).nodes)

	ans = graph.ContainsNode(*NewNode("two"))
	if !ans {
		t.Error("Error in ContainsNode")
	}

	ans = graph.ContainsNode(*NewNode("four"))
	if ans {
		t.Error("Error in ContainsNode")
	}
}

func TestPipeOperations(t *testing.T) {
	node1 := NewNode("node1")
	node2 := NewNode("node2")
	node3 := NewNode("node3")

	graph := NewGraph()
	graph.AddNode(*node1)
	graph.AddNode(*node2)
	graph.AddNode(*node3)

	pipe1 := NewPipe(*node1, *node2)
	pipe2 := NewPipe(*node1, *node3)

	ans := pipe1.equals(*pipe2)
	if ans {
		t.Error("Error in comparing non-equal pipes")
	}

	ans = pipe1.equals(*NewPipe(*NewNode("node1"), *NewNode("node2")))
	if !ans {
		t.Error("Error in comparing equal pipes")
	}

	ans = graph.ValidatePipe(*pipe1)
	if !ans {
		t.Error("Error in validating valid pipe")
	}

	ans = graph.ValidatePipe(*NewPipe(*NewNode("node1"), *NewNode("node1")))
	if ans {
		t.Error("Error in validating invalid pipe")
	}
}
