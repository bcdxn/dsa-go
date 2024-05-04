package ds

import "errors"

func Dbj2Hash(input string) uint {
	var hash uint = 5381

	for _, char := range input {
		hash = (hash << 5) + hash + uint(byte(char))
	}

	return hash
}

type HashTable struct {
	table []*List[string]
	len   int
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make([]*List[string], 10),
		len:   0,
	}
}

// Len returns the number of elements stored in the hashtable
func (ht HashTable) Len() int {
	return ht.len
}

// Len returns the capacity of the hash table
func (ht HashTable) Cap() int {
	return cap(ht.table)
}

// Add adds an element to the hash table. The HashTable will automatically resize if needed.
func (ht *HashTable) Add(s string) {
	hash := Dbj2Hash(s)
	capacity := len(ht.table)

	if (ht.Len() + 1) > (capacity / 2) {
		ht.resize()
		capacity = len(ht.table)
	}

	index := hash % uint(capacity)

	if ht.table[index] == nil {
		ht.table[index] = NewList[string]()
	}

	ht.table[index].AddHead(s)
	ht.len++
}

func (ht HashTable) Get(hash uint) (string, error) {
	index := hash % uint(len(ht.table))
	list := ht.table[int(index)]

	if list == nil {
		return "", errors.New("hash not found")
	}

	return list.Head.Elem, nil
}

func moveTableList(table []*List[string], list *List[string]) {
	// All items in the list have the same hash so we can calculate the hash using the first element
	// in the list
	hash := Dbj2Hash(list.Head.Elem)
	capacity := len(table)
	// Add the list at the proper index in the underlying slice
	index := hash % uint(capacity)
	table[index] = list
}

func (ht *HashTable) resize() {
	// Resize the underlying slice
	newTable := make([]*List[string], cap(ht.table)*2)
	// re-hash all items in the table
	for _, list := range ht.table {
		// Add all non-null lists to the new table
		if list != nil {
			moveTableList(newTable, list)
		}
	}
	// Update the underlying hash table slice to the newly allocated one
	ht.table = newTable
}
