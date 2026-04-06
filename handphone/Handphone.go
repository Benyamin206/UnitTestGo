package handphone

type Handphone struct {
	Merk         string
	Power_charge string
}

func (h Handphone) PlugInCharge() string {
	return "Charging Handphone " + h.Merk + " " + h.Power_charge + " watt"
}
