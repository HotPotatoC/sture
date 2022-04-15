# Sture

A collection of data structures based on Go 1.18+ Generics.

### Why this name?

I basically used an AI to generate the name **sture** 🗿. But... **sture** can be interpreted as "structure".

## Installation

```bash
go get github.com/HotPotatoC/sture
```

## Usage

Creating a priority queue using:

```go
import "github.com/HotPotatoC/sture/queue"

func main() {
    pq := queue.NewPriorityQueue[string]()

    pq.Enqueue("Adam", 1)
    pq.Enqueue("John", 3)
    pq.Enqueue("Bob", 2)

    top := pq.Peek()
    fmt.Println(top) // John
}
```

See more examples [here](./example/)

## Spec

Supported data structures at the moment:

- [Binary search tree](./tree/bstree.go)
- [AVL tree](./tree/avl_tree.go)
- [Linked List](./linkedlist/linked_list.go)
- [Queue](./queue/queue.go)
- [Priority Queue](./queue/priority_queue.go)
- [Ring Buffer](./queue/ring_buffer.go)
- [Deque](./queue/deque.go)
- [Stack](./stack/stack.go)
- [Hash Table](./hashtable/hashtable.go)
- [Chained Hash Table](./hashtable/chained_hashtable.go)
- more to come...

## Support

If this project is helpful to you, please consider supporting me by donating or just give this project a 🌟

<a href="https://www.buymeacoffee.com/hotpotato" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
