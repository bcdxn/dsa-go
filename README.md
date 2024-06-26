# DSA Practice in Go

## Data Structures

- [x] Linked List
- [x] Doubly Linked List
- [x] Queue (FIFO)
- [x] Stack (LIFO)
- [x] Hash Table + Collision Handling
- [x] BST
- [x] Heap
- [x] AVL
- [x] Graph
    - [x] Matrix
    - [x] Adjacency List

## Algorithms

### Searching

- [ ] Binary Search
- [ ] Linear Search

### Sorting

- [x] Bubble Sort
- [x] Insertion Sort
- [x] Merge Sort
- [x] Quick Sort
- [x] Counting Sort
- [ ] Shell Sort
- [ ] Heap Sort

### String Matching

### Tree Traversals

- [ ] Pre-order
- [ ] In-order
- [ ] Post-order
- [ ] Level-Order

### Graph Traversals

- [ ] Tree Height
- [ ] Detecting Cycles
- [ ] Breadth First
- [x] Depth First
- [ ] Dijkstra's algorithm (shortest path)

## Usage

Create a new Go module (if you don't already have existing module):

```bash
go mod init <your-module>
```

Install the `dsa-go` module:

```bash
go get -u github.com/bcdxn/dsa-go
```

## Tests

```bash
go test -v -coverprofile cover.out ./...
go tool cover -html cover.out -o cover.html
open cover.html
```
