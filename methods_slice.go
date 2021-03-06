package set

func (d *Datum) Slice() (values []interface{}) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	l := len(d.store)
	values = make([]interface{}, 0, l)

	for l--; l >= 0; l-- {
		values = append(values, d.store[l].Value)
	}

	return
}
