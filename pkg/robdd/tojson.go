package robdd

import (
	"encoding/json"
	"os"
)

type JsonIntermediate struct {
	// Nodes    []ROBDDNode    `json:"Nodes"`
	Root     int            `json:"Root"`
	Order    map[string]int `json:"Order"`
	RevOrder map[int]string `json:"ReverseOrder"`
	Cache    map[string]int `json:"Cache"`
	AllNodes []ROBDDNode    `json:"AllNodes"`
}

func BddToIntermediate(rbd *ROBDD, root int, order map[string]int) JsonIntermediate {
	result := JsonIntermediate{}
	// result.Nodes = make([]ROBDDNode, 0)

	// rbd.BFSRoot(root, &result)
	result.Root = root

	result.AllNodes = rbd.Nodes
	result.RevOrder = rbd.RevOrder
	result.Order = order
	result.Cache = rbd.Cache

	return result
}

func DumpIntermediate(intermediate *JsonIntermediate) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(intermediate)
}
