package internal_test

import (
	"fmt"

	"github.com/rwxrob/learnvim/internal"
)

func ExampleHas() {
	fmt.Println(internal.Has(`sh`, `bash`))
	fmt.Println(internal.Has(`python`))
	// Output:
	// true
	// false
}
