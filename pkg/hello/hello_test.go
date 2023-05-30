package hello_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/SaimonWoidig/golang-training/pkg/hello"
)

func TestSayHello(t *testing.T) {
	expected := "Hello Bob!"
	actual := hello.SayHello("Bob")

	if expected != actual {
		t.Errorf(`Hello("Bob") = "%s"; expected "%s"`, actual, expected)
	}
}

func TestSayHello10000(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	for i := 0; i < 10000; i++ {
		expected := fmt.Sprintf("Hello Bob %s!", strconv.Itoa(i))
		actual := hello.SayHello("Bob " + strconv.Itoa(i))

		if expected != actual {
			t.Errorf(`Hello("Bob") = "%s"; expected "%s"`, actual, expected)
		}
	}
}

func TestPrintSayHello(t *testing.T) {
	hello.PrintSayHello("Bob")
}
