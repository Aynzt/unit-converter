package length

func MeterToFeet(m Meter) Feet {
	return Feet(m * 3.28)
}

func FeetToMeter(f Feet) Meter {
	return Meter(f / 3.28)
}
