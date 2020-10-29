func init() {debug.SetGCPercent(-1)}

type Trie struct {
    isEnd bool
    children []*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        children: make([]*Trie, 26),
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    l := len(word)
    curNode := this
    for i := 0; i < l; i++ {
        ch := word[i] - 'a'
        if curNode.children[ch] == nil {
            curNode.children[ch] = &Trie {
                children: make([]*Trie, 26),
            }
        }
        curNode = curNode.children[ch]
    }
    curNode.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    l := len(word)
    curNode := this
    for i := 0; i < l; i++ {
        ch := word[i] - 'a'
        if curNode.children[ch] == nil {
            return false
        }
        curNode = curNode.children[ch]
    }
    return curNode.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    l := len(prefix)
    curNode := this
    for i := 0; i < l; i++ {
        ch := prefix[i] - 'a'
        if curNode.children[ch] == nil {
            return false
        }
        curNode = curNode.children[ch]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */