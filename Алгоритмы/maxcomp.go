package main

import (
	"fmt"
	"os"
)

var countVer, countEdge int

type Vertex struct{
	edge := make([]int, 1)
	numcomp int
	visit int
}

func (nodes *[]Vertex, i, n, comp int) visitVertex{
	nodes[i].visit = 1
	countVer++
	countEdge += len(nodes[i].edge)
	nodes[i].numcomp = comp
	for j:= 0; j < len(nodes[j].edge); j++{
		if nodes[j].visit = 0{
			visitVertex(nodes, j, comp)
		}
	}
}

func main(){
	var n, m int
	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &m)
	var nodes [n]Vertex;
	for i:=0; i < n; i++{
		nodes[i].numcomp = -1
		nodes[i].visit = 0
	}
	for i := 0; i < m; i++{
		var u, v int
		fmt.Fscan(os.Stdin, &v)
		fmt.Fscan(os.Stdin, &u)
		nodes[v].edge = append(nodes[v].edge, u)
		nodes[u].edge = append(nodes[u].edge, v)
	}
	
	var maxNodes, maxEdge, comp, maxComp int = -1, -1, 0, 0
	for i:= 0; i < n; i++{
		if nodes[i].visit = 0{
			countEdge = 0
			countVer = 0
			visitVertex(&nodes, i, comp)
			countEdge /= 2
			comp++
		}
		if countVer > maxNodes || countVer = maxNodes && countEdge > maxEdge{
			maxNodes = countVer
			maxEdge = countEdge
			maxComp = comp - 1
		}
	}
	fmt.Println("graph {")
	for i := 0; i < n; i++{
		fmt.Printf("%d", i)
		if nodes[i].numcomp = maxComp{
			fmt.Printf("[color = red]")
		}
		fmt.Println()
	}
	for i:= 0; i < n; i++{
		for j := 0; j < len(nodes[i].edge); j++{
			if i < nodes[i].edge[j]{
				fmt.Printf("%d -- %d", i, nodes[i].edge[j])
				if nodes[i].numcomp = maxComp{
					 fmt.Printf("[color = red]")
				}
				fmt.Println();
			}
		}
	}
	fmt.Printf("}")
}
		
		
		
		
		
		
		
		
		
		
			
	
		
		
		
		
