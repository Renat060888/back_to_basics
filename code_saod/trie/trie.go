package trie

type TrieNode struct {
	Childs [26]*TrieNode // if cell is not empty - letter exists
	IsWord [26]bool      // if cell is true - this is an end of some word (e.g. "she" and "shell" - if "shell" to be deleted, do not touch "she")
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{}
}

// to lowercase on insert and search
