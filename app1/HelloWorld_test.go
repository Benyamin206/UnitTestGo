package app1

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	value := HelloWorld("Andi")

	if value != "Hi Andi" {
		t.Error("Result must be Hi Andi but the value is " + value)
	}

	fmt.Println("Testing Done")
}

func TestHelloWorldBen(t *testing.T) {
	value := HelloWorldBen("Sbrn")
	if value != "Hi Ben Sbrn" {
		t.Fatal("The value is " + value + " but the expected is Hi Ben Sbrn")
	}
	fmt.Println("Testing Done")
}

// fail => menggagalkan unit test tetapi tetap melanjutkan program selanjutnya
// failNow => - II - tidak melanjutkan next program
// error => menggagalkan unit test dan memanggil fail dan dengan pesan
// fatal => menggagalkan unit test dan panggil failnow dan dengan pesan
