package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("Expected '%s' but got '%s'", expected, repeated)
	}
}

func Benchmarkrepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	//Output: aaaaa
}