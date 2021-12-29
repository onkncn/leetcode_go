package main

type Trie struct {
	End  bool
	Next []*Trie
}

func Constructor() Trie {
	root := Trie{false, make([]*Trie, 26)}
	return root
}

func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		if this.Next[word[i]-'a'] == nil {
			this.Next[word[i]-'a'] = &Trie{false, make([]*Trie, 26)}
		}
		this = this.Next[word[i]-'a']
	}
	this.End = true
}

func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if this.Next[word[i]-'a'] == nil {
			return false
		}
		this = this.Next[word[i]-'a']
	}
	return this.End
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if this.Next[prefix[i]-'a'] == nil {
			return false
		}
		this = this.Next[prefix[i]-'a']
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
func main() {

}
