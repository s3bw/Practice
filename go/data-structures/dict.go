/* Go maps with additional methods
- Set
- Delete
- Has
- Get
- Clear
- Size
- Keys
- Values
*/

package dictionary

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Key is a generic type allowing substitution
type Key generic.Type

// Value is a generic type allowing substitution
type Value generic.Type

// Dictionary the set of Items
type Dictionary struct {
	items map[Key]Value
	lock  sync.RWMutex
}

// Set adds a new item to the dictionary
func (d *Dictionary) Set(k Key, v Value) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.items == nil {
		d.items = make(map[Key]Value)
	}
	d.items[k] = v
}

// Delete a value from a Dictionary, given key
func (d *Dictionary) Delete(k Key) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	_, ok := d.items[k]
	if ok {
		delete(d.items, k)
	}
	return ok
}

// Has returns true if the key exists in the dictionary
func (d *Dictionary) Has(k Key) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()
	_, ok := d.items[k]
	return ok
}

// Get returns the value associated with the key
func (d *Dictionary) Get(k Key) Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.items[k]
}

// Clear removes all the items from the dictionary
func (d *Dictionary) Clear() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.items = make(map[Key]Value)
}

// Size returns the amount of elements in the dictionary
func (d *Dictionary) Size() int {
	d.lock.Lock()
	defer d.lock.Unlock()
	return len(d.items)
}

// Keys returns a slice of all the keys present
func (d *Dictionary) Keys() []Key {
	d.lock.RLock()
	defer d.lock.RUnlock()
	keys := []Key{}
	for i := range d.items {
		keys = append(keys, i)
	}
	return keys
}

// Values returns a slice of all the values present
func (d *Dictionary) Values() []Value {
	d.lock.RLock()
	defer d.lock.RUnlock()
	values := []Value{}
	for i := range d.items {
		values = append(values, d.items[i])
	}
	return values
}
