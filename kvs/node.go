package kvs

import (
)

type Node struct {
    Dir map[string]*Node
    Value string
}

func NewNode() *Node {
    node := &Node{
        Dir: make(map[string]*Node),
        Value: "/",
    }

    return node
}

// returns value at key
func (node *Node) getValue(key string) string {
    return node.Dir[key].Value
}

// deletes value at key
func (node *Node) delValue(key string) {
    delete(node.Dir, key)
}

// adds value at position specified by key
func (node *Node) putValue(key string, val string) {
    node.Dir[key] = NewNode()
    node.Dir[key].Value = val
}
