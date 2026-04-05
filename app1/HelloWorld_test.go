package app1

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestHelloWorld(t *testing.T) {
// 	value := HelloWorld("Andi")

// 	if value != "Hi Andi" {
// 		t.Error("Result must be Hi Andi but the value is " + value)
// 	}

// 	fmt.Println("Testing Done")
// }

func TestHelloWorldBen(t *testing.T) {
	value := HelloWorldBen("Sbrn")
	assert.Equal(t, "Hi Ben Sbrn", value, "Should be same")
	fmt.Println("Testing 1 Selesai")
}

func TestHelloCandraRequire(t *testing.T) {
	value := HelloCandra("Candra")

	require.Equal(t, "Candras", value, "Should be equal")
	fmt.Println("Selesaiii")
}

func TestHelloWorldName(t *testing.T) {
	value := HelloWorldName("Eko")

	assert.Equal(t, "Eko", value, "they should be equal")

	fmt.Println("Testing 2 Selesai")
}

func SkipTest(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Test tidak dapat berjalan di windows")
	}
	fmt.Println("Doneeee")
}

// fail => menggagalkan unit test tetapi tetap melanjutkan program selanjutnya
// failNow => - II - tidak melanjutkan next program
// error => menggagalkan unit test dan memanggil fail dan dengan pesan
// fatal => menggagalkan unit test dan panggil failnow dan dengan pesan
