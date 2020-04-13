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

坑点：
	// 一个9064767行的文件 每行做hashring，结果不一致 原因hashBytes只取4位 529和83 一致  ，这个库是一对多的关系，非多对一
	// node=1000 vitualSpots=20000
	//9063340 key=49f57cf0 node=529
	//9063340 key=49f57cf0 node=83