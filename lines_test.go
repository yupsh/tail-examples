package tail_test

import (
	"strings"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tail"
)

func ExampleTail_lines() {
	// echo "1\n2\n3\n4\n5" | tail -n 3
	gloo.MustRun(
		Tail(Lines(3), strings.NewReader("1\n2\n3\n4\n5")),
	)
	// Output:
	// 3
	// 4
	// 5
}
