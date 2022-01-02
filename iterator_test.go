package gopatterns

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	i := Iterator{
		tasks: []string{"one", "two", "three"},
	}
	for _, val, more := i.Next(); more; _, val, more = i.Next() {
		fmt.Println(val)
	}
}
