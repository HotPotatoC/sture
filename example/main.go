package main

import (
	"fmt"

	"github.com/HotPotatoC/sture/queue"
)

func main() {
    pq := queue.NewPriorityQueue[string]()

    pq.Enqueue("Adam", 1)
    pq.Enqueue("John", 3)
    pq.Enqueue("Bob", 2)

    top := pq.Peek()
    fmt.Println(top) // John
}