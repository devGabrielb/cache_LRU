package main

type Item struct {
	value any
	next  *Item
	prev  *Item
	key   string
}
type CacheLRU struct {
	head     *Item
	end      *Item
	values   map[string]*Item
	capacity int
}

func New(capacity int) *CacheLRU {
	return &CacheLRU{
		capacity: capacity,
	}
}

func (l *CacheLRU) Get(key string) any {
	data, ok := l.values[key]
	if !ok {
		return -1
	}
	if data == l.head {
		return data.value
	}

	if data.prev != nil {
		data.prev.next = data.next
	}
	if data.next != nil {
		data.next.prev = data.prev
	}
	if data == l.head {
		l.head = data.next
	}

	l.end.next = data
	data.prev = l.end
	l.end = data

	return data.value
}

func (l *CacheLRU) Set(key string, value any) {
	if len(l.values) == l.capacity {
		oldItem := l.head

		delete(l.values, oldItem.key)
		l.head.key = key
		l.head.value = value
		l.values[key] = l.head

	}
	newData := &Item{
		value: value,
		key:   key,
	}

	newData.next = l.head
	if l.head != nil {
		l.head.prev = newData
	}
	l.head = newData
	l.values[key] = newData
	if l.head.next == nil {
		l.end = l.head
	}
}

func main() {
}
