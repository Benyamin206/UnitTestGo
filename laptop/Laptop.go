package laptop

type Laptop struct {
	Merk         string
	Power_charge string
}

func (l Laptop) PlugInCharge() string {
	return "Charging Laptop " + l.Merk + " " + l.Power_charge + " watt"
}
