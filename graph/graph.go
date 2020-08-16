package graph

import (
	"bytes"
	"container/list"
	"log"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

type Node struct {
	name      string
	operation func()
}

func (n Node) Operation() func() {
	return n.operation
}

func (n Node) SetOperation(operation func()) {
	n.operation = operation
}

func (n Node) Name() string {
	return n.name
}

func (n Node) equals(node Node) bool {
	return n.name == node.name
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

type Pipe struct {
	source Node
	sink   Node
}

func (p Pipe) Sink() Node {
	return p.sink
}

func (p Pipe) Source() Node {
	return p.source
}

func NewPipe(source Node, sink Node) *Pipe {
	return &Pipe{source: source, sink: sink}
}

func (p Pipe) equals(pipe Pipe) bool {
	return p.source.equals(pipe.source) && p.sink.equals(pipe.sink)
}

type Graph struct {
	nodes list.List
	pipes list.List
}

func (g *Graph) Pipes() list.List {
	return g.pipes
}

func (g *Graph) Nodes() list.List {
	return g.nodes
}

func NewGraph() *Graph {
	return &Graph{
		nodes: list.List{},
		pipes: list.List{},
	}
}

func (g *Graph) AddNode(node Node) bool {
	if g.ContainsNode(node) {
		logger.Printf("Node %+v already exists", node)
		return false
	}
	g.nodes.PushBack(node)
	return true
}

func (g *Graph) RemoveNode(node Node) interface{} {
	e := g.nodes.Remove(&list.Element{Value: node})
	for p := g.pipes.Front(); p != nil; p.Next() {
		pipe, ok := p.Value.(Pipe)
		if ok && (pipe.source.equals(node) || pipe.sink.equals(node)) {
			g.pipes.Remove(p)
		}
	}
	return e
}

func (g *Graph) AddPipe(pipe Pipe) bool {
	if !g.ValidatePipe(pipe) {
		logger.Printf("Pipe %+v is not valid for graph %+v", pipe, g)
		return false
	}
	g.pipes.PushBack(pipe)
	return true
}

func (g *Graph) RemovePipe(pipe Pipe) interface{} {
	return g.pipes.Remove(&list.Element{Value: pipe})
}

func (g *Graph) ValidatePipe(pipe Pipe) bool {
	return !(&pipe.source == nil || &pipe.sink == nil || pipe.source.equals(pipe.sink)) && g.ContainsNode(pipe.source) &&
		g.ContainsNode(pipe.sink) && !g.isCyclicWith(pipe)
}

func (g *Graph) isCyclicWith(pipe Pipe) bool {
	workGraph := *NewGraph()
	workGraph.nodes.PushBackList(&g.nodes)
	workGraph.pipes.PushBackList(&g.pipes)
	// Do not call Graph.AddPipe(pipe) which calls this function to verify the pipe and hence causes an endless loop
	workGraph.pipes.PushBack(pipe)
	return workGraph.isCyclic()
}

func (g *Graph) isCyclic() bool {
	if &g == nil || g.nodes.Len() < 2 || g.pipes.Len() < 2 {
		return false
	}
	workGraph := NewGraph()
	workGraph.nodes.PushBackList(&g.nodes)
	workGraph.pipes.PushBackList(&g.pipes)
	roots := workGraph.GetRoots()
	for roots.Len() == 0 || workGraph.nodes.Len() == 0 {
		for n := roots.Front(); n != nil; n.Next() {
			node, ok := n.Value.(Node)
			if ok {
				workGraph.RemoveNode(node)
			}
		}
	}
	return workGraph.nodes.Len() == 0
}

func (g *Graph) ContainsNode(node Node) bool {
	var nodes = g.nodes
	for e := nodes.Front(); e != nil; e = e.Next() {
		n, ok := e.Value.(Node)
		if ok && n.equals(node) {
			return true
		}
	}
	return false
}

func (g *Graph) GetRoots() list.List {
	roots := *list.New()
	for n := g.nodes.Front(); n != nil; n.Next() {
		node, okNode := n.Value.(Node)
		if okNode {
			isRoot := true
			for p := g.pipes.Front(); p != nil; p.Next() {
				pipe, okPipe := p.Value.(Pipe)
				if okPipe && pipe.sink.equals(node) {
					isRoot = false
					break
				}
			}
			if isRoot {
				roots.PushBack(node)
			}
		}
	}
	return roots
}

func (g *Graph) GetLeaves() list.List {
	roots := *list.New()
	for n := g.nodes.Front(); n != nil; n.Next() {
		node, okNode := n.Value.(Node)
		if okNode {
			isRoot := true
			for p := g.pipes.Front(); p != nil; p.Next() {
				pipe, okPipe := p.Value.(Pipe)
				if okPipe && pipe.source.equals(node) {
					isRoot = false
					break
				}
			}
			if isRoot {
				roots.PushBack(node)
			}
		}
	}
	return roots
}
