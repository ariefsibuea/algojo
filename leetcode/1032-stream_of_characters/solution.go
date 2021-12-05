package streamofcharacters

// What to Do:
// 	Design an algorithm that accepts a stream of characters and checks if a suffix of these characters is a string of a given array of strings words.

// Input:
// 	["StreamChecker", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query"]
// 	[[["cd", "f", "kl"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["g"], ["h"], ["i"], ["j"], ["k"], ["l"]]
// Output:
// 	[null, false, false, false, true, false, true, false, false, false, false, false, true]

// Explanation:
// 	StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
// 	streamChecker.query("a"); // return False
// 	streamChecker.query("b"); // return False
// 	streamChecker.query("c"); // return False
// 	streamChecker.query("d"); // return True, because 'cd' is in the wordlist
// 	streamChecker.query("e"); // return False
// 	streamChecker.query("f"); // return True, because 'f' is in the wordlist
// 	streamChecker.query("g"); // return False
// 	streamChecker.query("h"); // return False
// 	streamChecker.query("i"); // return False
// 	streamChecker.query("j"); // return False
// 	streamChecker.query("k"); // return False
// 	streamChecker.query("l"); // return True, because 'kl' is in the wordlist

// How to Solve:
//  Implement trie, put the words into trie. Source: https://www.hackerearth.com/practice/data-structures/advanced-data-structures/trie-keyword-tree/tutorial/

type StreamChecker struct {
	Root   *Node
	Stream []byte
}

type Node struct {
	Next      map[byte]*Node
	IsEndNode bool
}

func NewNode() *Node {
	return &Node{
		Next:      map[byte]*Node{},
		IsEndNode: false,
	}
}

func Constructor(words []string) StreamChecker {
	streamChecker := StreamChecker{
		Root:   NewNode(),
		Stream: make([]byte, 0),
	}

	for _, word := range words {
		// reverse the order of letters because we check suffix
		currentNode, wordBytes := streamChecker.Root, []byte(word)
		for i := len(wordBytes) - 1; i >= 0; i-- {
			letter := wordBytes[i]
			if currentNode.Next[letter] == nil {
				currentNode.Next[letter] = NewNode()
			}
			currentNode = currentNode.Next[letter]
		}
		currentNode.IsEndNode = true
	}

	return streamChecker
}

func (this *StreamChecker) Query(letter byte) bool {
	this.Stream = append(this.Stream, letter)

	// because we want to check suffix, iterate the stream from end of stream
	currentNode := this.Root
	for i := len(this.Stream) - 1; i >= 0; i-- {
		letter = this.Stream[i]
		if currentNode.Next[letter] == nil {
			return false
		}

		currentNode = currentNode.Next[letter]
		if currentNode.IsEndNode {
			return true
		}
	}

	return currentNode.IsEndNode
}
