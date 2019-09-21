package pqueue

type Entry interface {
	// Compare compare with another entry, return true if a < b
	Less(a Entry) bool
	Value() interface{}
}
