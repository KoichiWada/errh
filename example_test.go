package errh_test

import (
	"fmt"

	"github.com/KoichiWada/errh"
)

func ExampleErrorf() {
	err := errh.Errorf("something wrong")
	fmt.Println(err)
	// Output: something wrong[example_test.go:10]
}
