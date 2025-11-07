package tail_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tail"
)

// This example demonstrates reading from a file instead of strings.NewReader
func ExampleTail_fromFile_basic() {
	// tail testdata/numbers.txt
	yup.MustRun(
		Tail(yup.File("testdata/numbers.txt")),
	)
	// Output:
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
	// 12
}

