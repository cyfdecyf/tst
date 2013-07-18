/*
Package tst is a pure Go implementation of *Ternary Search Trie*.

It's also called *Ternary Search Tree*, here's a detailed
description by the inventors Jon Bently and Robert Sedgwick,
http://www.drdobbs.com/database/ternary-search-trees/184410528.

It can be used as an efficient symbol table (as efficient as hash table), and
supports order traversal and many advanced search operations like prefix and
wildcard searches.
*/
package tst

// import "fmt"

type node struct {
	c     byte
	left  *node
	right *node
	mid   *node
	val   interface{}
}

// Trie represents a Ternary Search Trie. An empty value can be used to insert
// key/value pairs.
type Trie struct {
	root *node
	n    int
}

// Size returns number of elements in the trie.
func (t *Trie) Size() int {
	return t.n
}

// Put inserts a key/value pair into the trie. If the key is already in the
// trie, it's value will be updated. Note: empty strings will be ignored.
func (t *Trie) Put(key string, value interface{}) {
	if len(key) < 1 {
		return
	}

	t.n++
	pnd := &t.root
	i := 0
	c := key[0]
	for {
		if *pnd == nil {
			// fmt.Printf("new node: %c\n", c)
			*pnd = &node{c: c}
		}
		switch {
		case c < (*pnd).c:
			pnd = &(*pnd).left
		case c > (*pnd).c:
			pnd = &(*pnd).right
		case i < len(key)-1:
			pnd = &(*pnd).mid
			i++
			c = key[i]
		default:
			(*pnd).val = value
			// fmt.Println("add value:", value)
			return
		}
	}
}

// Get returns the value associated with key.
func (t *Trie) Get(key string) interface{} {
	if len(key) < 1 {
		return nil
	}

	nd := t.root
	i := 0
	c := key[0]
	// go down the tree
	for nd != nil {
		switch {
		case c < nd.c:
			nd = nd.left
		case c > nd.c:
			nd = nd.right
		case i < len(key)-1:
			nd = nd.mid
			i++
			c = key[i]
		default:
			return nd.val
		}
	}
	return nil
}

// GetShortestPrefix searches for the shortest key which is a prefix of the
// given key and returns the value. For example, if trie contains key "com"
// and "com.g", search for "com.google" will return value associated with
// "com".
func (t *Trie) GetShortestPrefix(key string) interface{} {
	// Most code copied from Get.
	if len(key) < 1 {
		return nil
	}

	nd := t.root
	i := 0
	c := key[0]
	for nd != nil {
		switch {
		case c < nd.c:
			nd = nd.left
		case c > nd.c:
			nd = nd.right
		// This is the only added code from Get.
		case nd.val != nil:
			return nd.val
		case i < len(key)-1:
			nd = nd.mid
			i++
			c = key[i]
		default:
			return nd.val
		}
	}
	return nil
}
