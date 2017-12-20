package kvs

import (
)

type Node struct {
    Dir map[string]*Node
    Value string
}

// constructor
func NewNode() *Node {
    node := &Node{
        Dir: make(map[string]*Node),
        Value: "/",
    }

    return node
}

// determines whether key exists in KVStore
func (node *Node) inStore(key string) bool {
    // assigns _ to Dir[key] or the empty value, OK gets a boolean
    // if statement evaluates OK
    if _, ok := node.Dir[key]; ok {
        return true
    }
    return false
}

// returns value at key
func (node *Node) getValue(key string) string {
    if node.inStore(key) {
        return node.Dir[key].Value
    }
    return ""
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
