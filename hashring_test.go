package hashring

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"testing"
	"time"
)

const (
	node1 = "192.168.1.1"
	node2 = "192.168.1.2"
	node3 = "192.168.1.3"
)

func getNodesCount(nodes nodesArray) (int, int, int) {
	node1Count := 0
	node2Count := 0
	node3Count := 0

	for _, node := range nodes {
		if node.nodeKey == node1 {
			node1Count += 1
		}
		if node.nodeKey == node2 {
			node2Count += 1
		}
		if node.nodeKey == node3 {
			node3Count += 1

		}
	}
	return node1Count, node2Count, node3Count
}

func TestHash(t *testing.T) {

	nodeWeight := make(map[string]int)
	nodeWeight[node1] = 2
	nodeWeight[node2] = 2
	nodeWeight[node3] = 3
	vitualSpots := 100

	hash := NewHashRing(vitualSpots)

	hash.AddNodes(nodeWeight)
	if hash.GetNode("1") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 := getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after AddNodes node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

	hash.RemoveNode(node3)
	if hash.GetNode("1") != node1 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node2 {
		t.Fatalf("expetcd %v got %v", node1, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after RemoveNode node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

	hash.AddNode(node3, 3)
	if hash.GetNode("1") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("1"))
	}
	if hash.GetNode("2") != node3 {
		t.Fatalf("expetcd %v got %v", node3, hash.GetNode("2"))
	}
	if hash.GetNode("3") != node2 {
		t.Fatalf("expetcd %v got %v", node2, hash.GetNode("3"))
	}
	c1, c2, c3 = getNodesCount(hash.nodes)
	t.Logf("len of nodes is %v after AddNode node1:%v, node2:%v, node3:%v", len(hash.nodes), c1, c2, c3)

}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	fmt.Printf("TimeTrack funcName:%s elapsed:%s \n", name, elapsed)
}

func TestSpeed(t *testing.T) {
	// node=1000 v=50000 on 16GB/win10/i7 init 35s
	defer TimeTrack(time.Now())
	nodeWeight := make(map[string]int)
	for i := 0; i < 1000; i += 1 {
		nodeWeight[strconv.Itoa(i)] = 1
	}
	vitualSpots := 50000
	hash := NewHashRing(vitualSpots)

	hash.AddNodes(nodeWeight)
}
