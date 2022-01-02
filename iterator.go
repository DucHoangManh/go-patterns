package gopatterns

type Iterator struct {
	tasks    []string
	position int
}

// Next will return the next task in the slice
// if there's more data to iterate, more will be true
func (t *Iterator) Next() (pos int, val string, more bool) {
	t.position++
	if t.position > len(t.tasks) {
		return t.position, "", false
	}
	return t.position, t.tasks[t.position-1], true
}
