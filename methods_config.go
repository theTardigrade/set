package set

func (d *Datum) copyConfig(d2 *Datum) {
	d.equalityTest = d2.equalityTest
	d.maximumValueCount = d2.maximumValueCount
	d.filter = d2.filter
	d.multiMode = d2.multiMode
}

func (d *Datum) SetConfig(c *Config) {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	d.setEqualityTest(c.EqualityTest)
	if m := c.MaximumValueCount; m != nil {
		d.setMaximumValueCount(*m)
	}
	d.setFilter(c.Filter)
	d.setMultiMode(c.MultiMode)
	if c := c.Capacity; c != nil {
		d.setCap(*c)
	}
}

// mutex should be locked before calling
func (d *Datum) setEqualityTest(equalityTest equalityTestFunc) (success bool) {
	if equalityTest == nil {
		return
	}

	d.equalityTest = equalityTest
	d.removeDuplicates()
	success = true

	return
}

func (d *Datum) SetEqualityTest(equalityTest equalityTestFunc) bool {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.setEqualityTest(equalityTest)
}

// mutex should be locked before calling
func (d *Datum) setMaximumValueCount(n int) (success bool) {
	if n < 0 {
		return
	}

	d.maximumValueCount = &n

	if s := d.store; s != nil {
		if l := len(s); l > n {
			d.store = s[:n]
		}
	}

	success = true

	return
}

func (d *Datum) SetMaximumValueCount(n int) bool {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.setMaximumValueCount(n)
}

// mutex should be locked before calling
func (d *Datum) setFilter(f filterFunc) (success bool) {
	if f == nil {
		return
	}

	d.filter = f

	if s := d.store; s != nil {
		if l := len(d.store); l > 0 {
			var modified bool

			for l--; l >= 0; l-- {
				if v := s[l].Value; !f(v) {
					d.removeOneFromIndex(l)
					modified = true
				}
			}

			if modified {
				d.clearCachedFields()
			}
		}
	}

	success = true

	return
}

func (d *Datum) SetFilter(f filterFunc) bool {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.setFilter(f)
}

// mutex should be locked before calling
func (d *Datum) setMultiMode(multiMode bool) (success bool) {
	if prevMultiMode := d.multiMode; multiMode != prevMultiMode {
		d.multiMode = multiMode
		success = true

		if prevMultiMode {
			for _, s := range d.store {
				s.Instances = 1
			}
		}
	}

	return
}

func (d *Datum) SetMultiMode(multiMode bool) bool {
	defer d.mutex.Unlock()
	d.mutex.Lock()

	return d.setMultiMode(multiMode)
}
