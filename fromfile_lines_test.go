package tail_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tail"
)

func ExampleTail_fromFile_lines() {
	// tail -n 3 testdata/numbers.txt
	yup.MustRun(
		Tail(Lines(3), yup.File("testdata/numbers.txt")),
	)
	// Output:
	// 10
	// 11
	// 12
}

