type TrieNode struct {
	Children map[byte]*TrieNode
	EndOfWord bool
}
type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: &TrieNode{
			Children: make(map[byte]*TrieNode),
		},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.root
	for ix := range word {
		if _, ok := curr.Children[word[ix]]; !ok {
			curr.Children[word[ix]] = &TrieNode{
				Children: make(map[byte]*TrieNode),
			}
		}
		curr = curr.Children[word[ix]]
	}
	curr.EndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(iy int, root *TrieNode) bool
	dfs = func(iy int, root *TrieNode) bool {
		curr := root

		for ix := iy; ix < len(word); ix++ {
			c := word[ix]
			if c == '.' {
				for _, child := range curr.Children {
					if dfs(ix+1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := curr.Children[c]; !ok {
					return false
				}
				curr = curr.Children[c]
			}
		}
		return curr.EndOfWord
	}
	return dfs(0, this.root)
}
