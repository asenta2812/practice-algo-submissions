type LRUCache struct {
    capacity int
	items    map[int]*list.Element
	queue    *list.List
}

type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
    if element, ok := this.items[key]; ok {
		// Nếu tồn tại, di chuyển nó lên đầu hàng đợi (Most Recently Used)
		this.queue.MoveToFront(element)
		return element.Value.(*entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if element, ok := this.items[key]; ok {
		// Nếu key đã tồn tại, cập nhật giá trị và đưa lên đầu
		this.queue.MoveToFront(element)
		element.Value.(*entry).value = value
		return
	}

	// Nếu vượt quá dung lượng, xóa phần tử ở cuối (Least Recently Used)
	if this.queue.Len() >= this.capacity {
		lru := this.queue.Back()
		if lru != nil {
			this.queue.Remove(lru)
			delete(this.items, lru.Value.(*entry).key)
		}
	}

	// Thêm phần tử mới vào đầu hàng đợi và lưu vào map
	newEntry := &entry{key, value}
	element := this.queue.PushFront(newEntry)
	this.items[key] = element
}
