package apply

import (
	"github.com/cheekybits/genny/generic"
)

// Domain is the domain of the mapping function
type Domain generic.Type

// Codomain is the codomain of the mapping function
type Codomain generic.Type

// DomainCodomainApplier is a generic implementation of the higher-order apply
// (a.k.a. map) function. It makes it possible to apply a function to each item
// in a list and return each result in a list.
type DomainCodomainApplier func(Domain) (Codomain, error)

// Apply takes each given item in a Domain list and returns the corresponding
// Codomain list of results from running the function wrapped by
// DomainCodomainApplier on each Domain item. If any error is returned, the
// whole process stops immediately.
func (a DomainCodomainApplier) Apply(dl []Domain) ([]Codomain, error) {
	cl := make([]Codomain, len(dl))
	for i, d := range dl {
		r, e := a(d)
		if e != nil {
			return nil, e
		}
		cl[i] = r
	}
	return cl, nil
}
