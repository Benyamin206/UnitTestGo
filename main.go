package main

import (
	"fmt"

	"github.com/Benyamin206/UnitTestGo/handphone"
	"github.com/Benyamin206/UnitTestGo/laptop"
	"github.com/Benyamin206/UnitTestGo/resource1"
)

func PrintCharge(r1 resource1.StopKontak) {
	fmt.Println(r1.PlugInCharge())
}

func main() {
	var mi1 handphone.Handphone = handphone.Handphone{
		Merk:         "Xiaomi 11 T",
		Power_charge: "90",
	}

	var laptop1 laptop.Laptop = laptop.Laptop{
		Merk:         "Acer",
		Power_charge: "200",
	}

	PrintCharge(mi1)
	PrintCharge(laptop1)
}
