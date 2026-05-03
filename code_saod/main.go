package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"saod/circular_buffer"
	"saod/heap"
	"saod/stack"
	"saod/tree"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type ContainerDrawer struct {
	outputName string
	links      map[string][]string
}

func (cd *ContainerDrawer) VisitNode(name string, childs []string) {

	if cd.links == nil {
		cd.links = make(map[string][]string)
	}

	if _, found := cd.links[name]; !found {
		cd.links[name] = childs
	}
}

func (cd *ContainerDrawer) Draw() {

	drawToFile := func(g graph.Graph[string, string]) {
		file, err := os.Create(cd.outputName)
		if err != nil {
			panic(err)
		}

		err = draw.DOT(g, file)
		if err != nil {
			panic(err)
		}

		err = file.Sync()
		if err != nil {
			panic(err)
		}

		err = file.Close()
		if err != nil {
			panic(err)
		}
		fmt.Printf("container was drawn to file: %s\n", cd.outputName)

		cmd := exec.Command("dot", "-Tpng", cd.outputName, "-o", "dot.png")
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		if err := cmd.Start(); err != nil {
			fmt.Println("failed to start: " + stderr.String())
			panic(err)
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println("failed to wait: " + stderr.String())
			panic(err)
		}
	}

	fmt.Printf("map for drawing: %v\n", cd.links)

	g := graph.New(graph.StringHash, graph.Directed())

	if len(cd.links) == 0 {
		fmt.Printf("WARN: no links in set, nothing to draw\n")
		drawToFile(g)
		return
	}

	// create verices
	for k, _ := range cd.links {

		var err error
		if strings.Contains(k, "r") {
			err = g.AddVertex(k, graph.VertexAttribute("color", "red"), graph.VertexAttribute("style", "filled"))
		} else {
			err = g.AddVertex(k)
		}

		if err != nil {
			fmt.Printf("failed to add vertex with value: %s\n", k)
		}
	}

	// create connections between them
	for k, v := range cd.links {
		for _, c := range v {

			fmt.Printf("draw edge between %s and %s\n", k, c)

			err := g.AddEdge(k, c)
			if err != nil {
				fmt.Printf("failed to add edge between values: %s and %s\n", k, c)
			}
		}
	}

	// output
	drawToFile(g)
}

type ContainerPrinter struct {
	data string
}

func (cp *ContainerPrinter) VisitNode(name string, childs []string) {

	cp.data += name + " "
}

func (cp *ContainerPrinter) Print() {

	fmt.Println(cp.data)
}

func (cp *ContainerPrinter) Clear() {

	cp.data = ""
}

func heapCheck() {

	// heap
	heap := heap.NewHeap(heap.HEAP_POLICY_MAX)

	heap.AddValue(10)
	heap.AddValue(15)
	heap.AddValue(20)
	heap.AddValue(30)

	heap.AddValue(40)
	heap.AddValue(50)

	heap.AddValue(100)

	heapDrawer := ContainerDrawer{outputName: "./heap.gv"}
	heap.AcceptVisitor(&heapDrawer)
	heapDrawer.Draw()
}

func redBlackTreeCheck() {

	// tree
	redBlackTree := tree.NewRBTree()

	redBlackTree.Insert(10)
	redBlackTree.Insert(15)
	redBlackTree.Insert(5)

	redBlackTree.Insert(20)
	redBlackTree.Insert(17)

	redBlackTree.Insert(3)

	redBlackTree.Insert(50)

	treePrinter := ContainerPrinter{}

	redBlackTree.AcceptVisitor(&treePrinter, tree.RB_TREE_TRAVERSE_PRE_ORDER)
	treePrinter.Print()
	treePrinter.Clear()

	redBlackTree.AcceptVisitor(&treePrinter, tree.RB_TREE_TRAVERSE_IN_ORDER)
	treePrinter.Print()
	treePrinter.Clear()

	redBlackTree.AcceptVisitor(&treePrinter, tree.RB_TREE_TRAVERSE_POST_ORDER)
	treePrinter.Print()
	treePrinter.Clear()

	treeDrawer := ContainerDrawer{outputName: "./tree.gv"}
	redBlackTree.AcceptVisitor(&treeDrawer, tree.RB_TREE_TRAVERSE_PRE_ORDER)
	treeDrawer.Draw()
}

func stackCheck() {

	// stack
	stack := stack.NewStack[int]()
	fmt.Printf("stack initial size: %d\n", stack.GetSize())

	stack.Push(30).Push(20).Push(10)

	fmt.Printf("stack size: %d\n", stack.GetSize())

	if val := stack.Pop(); val != nil {
		fmt.Printf("stack's value is: %d\n", *val)
	} else {
		fmt.Printf("stack's value is unknown\n")
	}

	if val := stack.Pop(); val != nil {
		fmt.Printf("stack's value is: %d\n", *val)
	} else {
		fmt.Printf("stack's value is unknown\n")
	}

	if val := stack.Pop(); val != nil {
		fmt.Printf("stack's value is: %d\n", *val)
	} else {
		fmt.Printf("stack's value is unknown\n")
	}

	if val := stack.Pop(); val != nil {
		fmt.Printf("stack's value is: %d\n", *val)
	} else {
		fmt.Printf("stack's value is unknown\n")
	}
}

func bTreeCheck() {

	if len(os.Args) > 1 && os.Args[1] == "-i" {

		bTree := tree.NewBTree()
		var prevBTree *tree.BTree
		insert := true

		for true {
			var cmd string
			_, err := fmt.Scan(&cmd)
			if err != nil {
				fmt.Printf("failed to read value, try again (reason: %s)\n", err.Error())
				continue
			}

			if cmd == "i" {
				insert = true
				fmt.Printf("insert mode\n")
			} else if cmd == "d" {
				insert = false
				fmt.Printf("delete mode\n")
			} else if cmd == "p" {
				bTree = prevBTree

				bTreeDrawer := ContainerDrawer{outputName: "./btree.gv"}
				bTree.AcceptVisitor(&bTreeDrawer)
				bTreeDrawer.Draw()
			} else if strings.HasPrefix(cmd, "s") {
				// save current sequence of inputs to file
			} else if strings.HasPrefix(cmd, "l") {
				// load sequence of inputs from file
			} else if cmd == "q" {
				insert = false
				fmt.Printf("exit, bye! :)\n")
				os.Exit(0)
			} else {
				value, err := strconv.Atoi(cmd)
				if err != nil {
					fmt.Printf("failed to convert value to number, try again (reason: %s)\n", err.Error())
					continue
				}

				if insert {
					prevBTree = bTree.Clone()
					bTree.Insert(value)
				} else {
					prevBTree = bTree.Clone()
					bTree.RemoveTry2(value)
				}

				bTreeDrawer := ContainerDrawer{outputName: "./btree.gv"}
				bTree.AcceptVisitor(&bTreeDrawer)
				bTreeDrawer.Draw()
			}
		}
	}

	// b-tree
	fmt.Printf("// B-TREE ------------------------------------------------------------------------------------------\n")
	bTree := tree.NewBTree()
	fmt.Printf("b-tree size is %d\n", bTree.GetSize())

	bTree.Insert(10)
	bTree.Insert(20)
	bTree.Insert(30)
	fmt.Printf("b-tree size is %d\n", bTree.GetSize())

	bTree.Insert(5)
	bTree.Insert(15)

	bTree.Insert(35)
	bTree.Insert(40)

	bTree.Insert(45)

	//
	bTree.Insert(38)
	bTree.Insert(50)
	bTree.Insert(60)
	bTree.Insert(70)
	bTree.Insert(75)
	bTree.Insert(55)
	fmt.Printf("b-tree size is %d\n", bTree.GetSize())

	bTreeDrawer := ContainerDrawer{outputName: "./btree.gv"}
	bTree.AcceptVisitor(&bTreeDrawer)
	bTreeDrawer.Draw()
}

func rbTreeCheck() {

	if len(os.Args) > 1 && os.Args[1] == "-i" {

		bTree := tree.NewRBTree()
		var prevBTree *tree.RBTree
		insert := true

		for true {
			var cmd string
			_, err := fmt.Scan(&cmd)
			if err != nil {
				fmt.Printf("failed to read value, try again (reason: %s)\n", err.Error())
				continue
			}

			if cmd == "i" {
				insert = true
				fmt.Printf("insert mode\n")
			} else if cmd == "d" {
				insert = false
				fmt.Printf("delete mode\n")
			} else if cmd == "p" {
				bTree = prevBTree

				bTreeDrawer := ContainerDrawer{outputName: "./rbtree.gv"}
				bTree.AcceptVisitor(&bTreeDrawer, tree.RB_TREE_TRAVERSE_IN_ORDER)
				bTreeDrawer.Draw()
			} else if strings.HasPrefix(cmd, "s") {
				// save current sequence of inputs to file
			} else if strings.HasPrefix(cmd, "l") {
				// load sequence of inputs from file
			} else if cmd == "q" {
				insert = false
				fmt.Printf("exit, bye! :)\n")
				os.Exit(0)
			} else {
				value, err := strconv.Atoi(cmd)
				if err != nil {
					fmt.Printf("failed to convert value to number, try again (reason: %s)\n", err.Error())
					continue
				}

				if insert {
					prevBTree = bTree.Clone()
					bTree.Insert(value)
				} else {
					prevBTree = bTree.Clone()
					bTree.Remove(value)
				}

				bTreeDrawer := ContainerDrawer{outputName: "./rbtree.gv"}
				bTree.AcceptVisitor(&bTreeDrawer, tree.RB_TREE_TRAVERSE_IN_ORDER)
				bTreeDrawer.Draw()
			}
		}
	}

	// rb-tree
	fmt.Printf("// RB-TREE ------------------------------------------------------------------------------------------\n")
	bTree := tree.NewBTree()
	fmt.Printf("rb-tree size is %d\n", bTree.GetSize())

	bTree.Insert(20)
	bTree.Insert(30)
	bTree.Insert(10)
	fmt.Printf("rb-tree size is %d\n", bTree.GetSize())

	bTree.Insert(35)
	fmt.Printf("rb-tree size is %d\n", bTree.GetSize())

	bTreeDrawer := ContainerDrawer{outputName: "./rbtree.gv"}
	bTree.AcceptVisitor(&bTreeDrawer)
	bTreeDrawer.Draw()
}

// implemented by user code
type NetworkListener interface {
	OnEvent(request string, response *string)
}

// implemented by network entity
type NetworkClient interface {
	Make(request string) string
}

func Sortings() {

	// s := []int{5, 2}
	// s := []int{10, 3, 2, 1}
	// s := []int{2, 5, 4, 10, 3, 2, 1, 12}
	s := []int{}
	cycles := 11
	for cycles > 0 {
		s = append(s, rand.Intn(100))
		cycles--
	}
	fmt.Println(s)
	tree.SortSliceMerge(s)
	fmt.Println(s)
}

func main() {

	rbTreeCheck()
	return

	cb, err := circular_buffer.NewCircularBuffer(4)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("cr created!")
	}
	cb.Push2(1)
	cb.Push2(2)
	cb.Push2(3)
	cb.Push2(4)
	cb.Push2(5)
	cb.Push2(6)
	cb.Push2(7)
	cb.Push2(8)
	cb.Push2(9)

	fmt.Printf(">> size: %d\n", cb.GetSize())

	val := cb.Pop2()
	for val != nil {
		fmt.Println(*val)
		val = cb.Pop2()
	}

	fmt.Printf(">> size: %d\n", cb.GetSize())

	return

	fmt.Println("\xf0\x9f\x90\xb5")
	fmt.Println("\xf0\x9f\x99\x88")
	fmt.Println("\xf0\x9f\x99\x89")

	fmt.Printf("number of virtual cores (P) is %d, so we have %d OS threads (M)\n", runtime.NumCPU(), runtime.NumCPU())

	bTreeCheck()
}
