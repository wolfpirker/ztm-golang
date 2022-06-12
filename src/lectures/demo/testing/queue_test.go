// use `go test -v ./demo/testing` to run tests
package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		// since we didn't add anything to our queue q should be equal to 3 initially
		if len(q.items) != i {
			t.Errorf("incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("failed to append item %v to queue", i)
		}
	}
	// add this point our queue is full, we should not be able to add more things!
	if q.Append(4) {
		t.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order: %v, want: %v", item, i)
		}
	}
	item, ok := q.Next()
	if ok {
		// queue should already be empty here!
		t.Errorf("should not be any more items in queue, got: %v", item)
	}
}
