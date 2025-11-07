package tail_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tail"
)

func ExampleTail_basic() {
	// echo "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12" | tail
	yup.MustRun(
		Tail(strings.NewReader("1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12")),
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

