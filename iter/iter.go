package main
import "fmt"

func main() {
    numIter := numIterator{iterFunc: num, start: 0, limit: 2}
    fmt.Println(iterAll(&numIter))
}

type iterator interface {
    hasNext() bool
    getNext() []int
}

func iterAll(numIterator iterator) []int {
    result := []int{}
    for numIterator.hasNext() {
        nums := numIterator.getNext()
        result = append(result, nums...)
    }
    return result
}

type numIterator struct {
    iterFunc func(start, limit int) []int
    start int
    limit int
}

func (n *numIterator) hasNext() bool {
    if len(n.iterFunc(n.start, n.limit)) == 0 {
        return false
    }
    return true
}

func (n *numIterator) getNext() []int {
    result := n.iterFunc(n.start, n.limit)
    n.start += n.limit
    return result
}

func num(start int, limit int) []int {
    result := []int{}
    for i := start; i < start+limit; i++ {
        if i >= 10 {
            break
        }
        result = append(result, i)
    }
    return result
}
