package mapping_test

import (
	"fmt"
)

// eventually: go:generate genny -out=gen_mapping_test.go get mapping/mapping.go "Domain=int Codomain=string"
// but for now:
//go:generate genny -in=mapping.go -out=gen_mapping_test.go -pkg mapping_test gen "Domain=int Codomain=string"

func ExampleEvennessMapping() {
	evenness := IntToStringMapping(func(i int) (string, error) {
		if i%2 == 0 {
			return "even", nil
		}
		return "odd", nil
	})

	inputArray := []int{}
	for i := -2; i <= 2; i++ {
		inputArray = append(inputArray, i)
	}

	outputArray, err := evenness.MapAll(inputArray)
	if err != nil {
		panic(err)
	}

	fmt.Print(outputArray)
	// Output: [even odd even odd even]
}
