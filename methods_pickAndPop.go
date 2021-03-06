package set

import (
	"math/rand"
	"time"
)

func (d *Datum) Pop() (value interface{}) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	value, index := d.pick()
	if value != nil {
		d.removeOneFromIndex(index)
		d.clearCachedFields()
	}

	return
}

// mutex should be read-locked before calling
func (d *Datum) pick() (value interface{}, index int) {
	if l := len(d.store); l > 0 {
		index = rand.Intn(l)
		value = d.store[index]
	}

	return
}

func (d *Datum) Pick() (value interface{}) {
	defer d.mutex.RUnlock()
	d.mutex.RLock()

	value, _ = d.pick()

	return
}

// seed the random-number generator for use in the pick method
func init() {
	rand.Seed(time.Now().UnixNano())
}
