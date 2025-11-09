package tail_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tail"
)

func ExampleTail_fromFile_lines() {
	// tail -n 3 testdata/numbers.txt
	gloo.MustRun(
		Tail(Lines(3), gloo.File("testdata/numbers.txt")),
	)
	// Output:
	// 10
	// 11
	// 12
}
