# hashring

A golang consistent hashring

Install
===

	go get github.com/g4zhuj/hashring

Usage
===


```
// virtualSpots means virtual spots created by each node
nodeWeight := make(map[string]int)
nodeWeight["node1"] = 1
nodeWeight["node2"] = 1
nodeWeight["node3"] = 2
vitualSpots := 100
hash := NewHashRing(virtualSpots)
	
	
//add nodes
hash.AddNodes(nodeWeight)
	
//remove node
hash.RemoveNode("node3")

	
//add node
hash.AddNode("node3", 3)

	
//get key's node
node := hash.GetNode("key")

```