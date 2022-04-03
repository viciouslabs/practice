package main

import (
	"fmt"
	"strconv"
)

type UnionFind struct {
	root []int
}

func (uf *UnionFind) init(x int) {
	// uf.root = make([]int{}
	for i := 0; i < int(x); i++ {
		uf.root = append(uf.root, i)
	}
}

func (uf *UnionFind) printIt() {
	for i := 0; i < len(uf.root); i++ {
		fmt.Print(strconv.Itoa(uf.root[i]) + ", ")
	}
	fmt.Println("")
}

func (uf *UnionFind) find(x int) int {
	return uf.root[x]
}

func (uf *UnionFind) isConnected(x int, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *UnionFind) union(x int, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX != rootY {
		for i := 0; i < len(uf.root); i++ {
			if uf.root[i] == rootY {
				uf.root[i] = rootX
			}
		}
	}
}

func NewUnionFind(x int) UnionFind {
	uf := UnionFind{}
	uf.init(x)
	return uf
}

func main() {
	uf := NewUnionFind(10)
	// 1-2-5-6-7 3-8-9 4
	uf.union(1, 2)
	uf.printIt()
	uf.union(2, 5)
	uf.union(5, 6)
	uf.union(6, 7)
	uf.union(3, 8)
	uf.union(8, 9)
	fmt.Println(uf.isConnected(1, 5)) // true
	fmt.Println(uf.isConnected(5, 7)) // true
	fmt.Println(uf.isConnected(4, 9)) // false
	// 1-2-5-6-7 3-8-9-4
	uf.union(9, 4)
	fmt.Println(uf.isConnected(4, 9)) // true
}
