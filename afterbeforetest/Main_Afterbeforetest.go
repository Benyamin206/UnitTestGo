package afterbeforetest

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Before test")

	code := m.Run()

	log.Println("After test")

	os.Exit(code)
}
