package maps

import (
	"sync"

	"github.com/metabition/genny/generic"
)

type KeyType generic.Type
type ValueType generic.Type

// ConMapKeyTypeValueType is a concurrent safe map storing ValueType values,
// for KeyType keys.
type ConMapKeyTypeValueType interface {
	// Get the ValueType for the specified key.
	Get(key KeyType) ValueType
	// Get the ValueType for the specified key and a bool indicating
	// whether a value was found or not.
	GetOK(key KeyType) (ValueType, bool)
	// Set the ValueType for the specified key.
	Set(key KeyType, value ValueType)
	// Delete the ValueType for the specified key and return it.
	// Second argument returns true if the value was deleted.
	Delete(key KeyType) (ValueType, bool)
}

type conMapKeyTypeValueType struct {
	m map[KeyType]ValueType
	l sync.RWMutex
}

// make sure conMapKeyTypeValueType is a ConMapKeyTypeValueType
var _ ConMapKeyTypeValueType = (*conMapKeyTypeValueType)(nil)

// NewConMapKeyTypeValueType creates a new concurrent safe map storing ValueType values,
// for KeyType keys.
func NewConMapKeyTypeValueType() ConMapKeyTypeValueType {
	return &conMapKeyTypeValueType{
		m: make(map[KeyType]ValueType),
	}
}

func (cm *conMapKeyTypeValueType) Get(k KeyType) ValueType {
	cm.l.RLock()
	v := cm.m[k]
	cm.l.RUnlock()
	return v
}
func (cm *conMapKeyTypeValueType) GetOK(k KeyType) (ValueType, bool) {
	cm.l.RLock()
	v, ok := cm.m[k]
	cm.l.RUnlock()
	return v, ok
}
func (cm *conMapKeyTypeValueType) Set(k KeyType, v ValueType) {
	cm.l.Lock()
	cm.m[k] = v
	cm.l.Unlock()
}
func (cm *conMapKeyTypeValueType) Delete(k KeyType) (ValueType, bool) {
	cm.l.Lock()
	v, ok := cm.m[k]
	if ok {
		delete(cm.m, k)
	}
	cm.l.Unlock()
	return v, ok
}
