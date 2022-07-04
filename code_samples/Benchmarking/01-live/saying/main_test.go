package saying

import (
	"fmt"
	"testing"
)

func testGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear James" {
		t.Error("got", s, "expected", "Welcome my dear James")
	}
}

// The example code block actaully shows up as an example on the docs

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}