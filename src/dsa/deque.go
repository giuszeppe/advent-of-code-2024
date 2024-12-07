package dsa

type Deque[T any] struct {
	items []T
}

func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

func (d *Deque[T]) Len() int {
	return len(d.items)
}

func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.Len() == 0 {
		var noop T
		return noop, false
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.Len() == 0 {
		var noop T
		return noop, false
	}

	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}
