package maps_test

import (
	"testing"

	"github.com/metabition/gennylib/maps"
	"github.com/metabition/is"
)

func TestMap(t *testing.T) {
	is := is.New(t)

	k1 := new(maps.KeyType)
	k2 := new(maps.KeyType)
	v1 := new(maps.ValueType)
	v2 := new(maps.ValueType)

	m := maps.NewConMapKeyTypeValueType()
	_, ok := m.GetOK(k1)
	is.Equal(ok, false)

	m.Set(k1, v1)
	is.Equal(m.Get(k1), v1)
	m.Set(k2, v2)
	is.Equal(m.Get(k2), v2)

	v, ok := m.GetOK(k1)
	is.Equal(ok, true)
	is.Equal(v, v1)

	// delete k2
	v, ok = m.Delete(k2)
	is.Equal(ok, true)
	is.Equal(v, v2)

	_, ok = m.GetOK(k2)
	is.Equal(ok, false)

}
