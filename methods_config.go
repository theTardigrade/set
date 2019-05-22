package set

func (d *datum) SetEqualityTest(equalityTest equalityTestFunc) (success bool) {
	if equalityTest != nil {
		return
	}

	success = true

	defer d.mutex.Unlock()
	d.mutex.Lock()

	d.equalityTest = equalityTest

	return
}

// mutex should be locked before calling
func (d *datum) makeStore(capacity int) {
	d.store = make(storeData, 0, capacity)
}