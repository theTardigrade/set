package set

// mutex should be locked before calling
func (d *datum) clear(capacity int) {
	d.makeStore(capacity)
	d.clearCachedHash()
}

func (d *datum) Clear(keepCapacity bool) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	var capacity int

	if keepCapacity {
		capacity = cap(d.store)
	}

	d.clear(capacity)
}