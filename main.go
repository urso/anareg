package main

import (
	"flag"
	"fmt"
	"regexp/syntax"
	"strings"

	"github.com/tmc/dot"
)

type nodeSet struct {
	nodes []*node
}

type node struct {
	r    *syntax.Regexp
	node *dot.Node
}

var nodeID uint64

func main() {
	flag.Parse()
	pattern := strings.Join(flag.Args(), " ")

	r, err := syntax.Parse(pattern, syntax.Perl)
	if err != nil {
		fmt.Println(err)
		return
	}

	graph := dot.NewGraph("G")
	buildGraph(graph, r)
	fmt.Println(graph)
}

func buildGraph(g *dot.Graph, r *syntax.Regexp) {
	visited := newSet()
	root := newNode(r)
	g.AddNode(root.node)
	visited.add(root)
	doBuildGraph(g, visited, root)
}

func doBuildGraph(g *dot.Graph, visited *nodeSet, parent *node) {
	for _, r := range parent.r.Sub {
		found := true
		sub := visited.find(r)
		if sub == nil {
			found = false
			sub = newNode(r)
			g.AddNode(sub.node)
			visited.add(sub)
		}

		e := dot.NewEdge(parent.node, sub.node)
		g.AddEdge(e)

		if found {
			continue
		}

		doBuildGraph(g, visited, sub)
	}
}

var opNames = map[syntax.Op]string{
	syntax.OpNoMatch:        "OpNoMatch",
	syntax.OpEmptyMatch:     "OpEmptyMatch",
	syntax.OpLiteral:        "OpLiteral",
	syntax.OpCharClass:      "OpCharClass",
	syntax.OpAnyCharNotNL:   "OpAnyCharNotNL",
	syntax.OpAnyChar:        "OpAnyChar",
	syntax.OpBeginLine:      "OpBeginLine",
	syntax.OpEndLine:        "OpEndLine",
	syntax.OpBeginText:      "OpBeginText",
	syntax.OpEndText:        "OpEndText",
	syntax.OpWordBoundary:   "OpWordBoundary",
	syntax.OpNoWordBoundary: "OpNoWordBoundary",
	syntax.OpCapture:        "OpCapture",
	syntax.OpStar:           "OpStar",
	syntax.OpPlus:           "OpPlus",
	syntax.OpQuest:          "OpQuest",
	syntax.OpRepeat:         "OpRepeat",
	syntax.OpConcat:         "OpConcat",
	syntax.OpAlternate:      "OpAlternate",
}

func newNode(r *syntax.Regexp) *node {
	id := nodeID
	nodeID++

	label := fmt.Sprintf("%v: `%v`", opNames[r.Op], r)

	n := dot.NewNode(fmt.Sprintf("node%v", id))
	n.Set("label", label)
	return &node{
		node: n,
		r:    r,
	}
}

func newSet() *nodeSet {
	return &nodeSet{nodes: make([]*node, 0)}
}

func (s *nodeSet) find(r *syntax.Regexp) *node {
	for _, n := range s.nodes {
		if n.r == r {
			return n
		}
	}
	return nil
}

func (s *nodeSet) add(n *node) {
	s.nodes = append(s.nodes, n)
}
