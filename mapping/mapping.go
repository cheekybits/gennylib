package mapping

import (
	"github.com/cheekybits/genny/generic"
)

// Domain is the domain of the mapping function
type Domain generic.Type

// Codomain is the codomain of the mapping function
type Codomain generic.Type

// DomainToCodomainMapping is a generic implementation of the higher-order map
// (a.k.a. apply) function. It makes it possible to apply a function to each
// item in a list and return each result in a list.
type DomainToCodomainMapping func(Domain) (Codomain, error)

// MapAll takes each given item in a Domain list and returns the corresponding
// Codomain list of results from running the function wrapped by
// DomainCodomainApplier on each Domain item. If any error is returned, the
// whole process stops immediately.
func (m DomainToCodomainMapping) MapAll(dl []Domain) ([]Codomain, error) {
	cl := make([]Codomain, len(dl))
	for i, d := range dl {
		r, e := m(d)
		if e != nil {
			return nil, e
		}
		cl[i] = r
	}
	return cl, nil
}
