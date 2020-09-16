package main

import "fmt"

/*
实现一个 Trie (前缀树)，包含insert,search, 和startsWith这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母a-z构成的。
保证所有输入均为非空字符串。
*/

func main() {
	trie := NewTrie()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}

type Trie struct {
	KeyTag   bool           //节点所代表的字符串是否是完整的字符串
	Val      rune           //节点值，每个节点所包含的字符：考虑汉字的情况，这里用rune
	Children map[rune]*Trie //叶子节点，也可以用slice
}

/** Initialize your data structure here. */
func NewTrie() *Trie {
	t := &Trie{
		KeyTag:   false,
		Val:      -1,
		Children: make(map[rune]*Trie),
	}
	return t
}

func (t *Trie) init() {
	t.Children = make(map[rune]*Trie)
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	var node = t
	for i := 0; i < len(word); i++ {
		if node.Children == nil {
			node.init()
		}
		if next := node.Children[rune(word[i])]; next == nil { //如果不存在，则插入
			next = &Trie{
				KeyTag:   i == len(word)-1,
				Val:      rune(word[i]),
				Children: nil,
			}
			node.Children[rune(word[i])] = next
			node = next
		} else {
			if i == len(word)-1 {
				next.KeyTag = true
				break
			}
			node = next
		}
	}
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	var node = t
	for i := 0; i < len(word); i++ {
		if node.Children == nil {
			return false
		}
		if next := node.Children[rune(word[i])]; next == nil {
			return false
		} else {
			if i == len(word)-1 {
				return next.KeyTag
			}
			node = next
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartWith(prefix string) bool {
	var node = t
	for i := 0; i < len(prefix); i++ {
		if node.Children == nil {
			return false
		}
		if next := node.Children[rune(prefix[i])]; next == nil {
			return false
		} else {
			node = next
		}
	}
	return true
}
