package mapping

import (
	"errors"
	"testing"

	"github.com/cheekybits/is"
)

func TestMapAll(T *testing.T) {

	T.Run("happy path", func(t *testing.T) {
		input := []Domain{}
		for i := 0; i < 3; i++ {
			input = append(input, Domain(i))
		}

		data := make(map[Domain]Codomain, len(input))
		expected := make([]Codomain, len(input))
		for i, d := range input {
			data[d] = Codomain(i)
			expected[i] = data[d]
		}

		testMapper := func(d Domain) (Codomain, error) {
			r, present := data[d]
			if !present {
				return nil, nil
			}
			return r, nil
		}

		mapping := DomainToCodomainMapping(testMapper)
		actual, err := mapping.MapAll(input)

		isnt := is.New(t)
		isnt.NoErr(err)
		isnt.Equal(expected, actual)
	})

	T.Run("sad path", func(t *testing.T) {
		input := []Domain{}
		for i := 0; i < 3; i++ {
			input = append(input, Domain(i))
		}

		testMapper := func(d Domain) (Codomain, error) {
			return nil, errors.New("whatever")
		}

		mapping := DomainToCodomainMapping(testMapper)
		actual, err := mapping.MapAll(input)

		isnt := is.New(t)
		isnt.Err(err)
		isnt.Nil(actual)
	})

}
