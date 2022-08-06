package weight

func PoundtoKg(p Pound) Kilogram {
	return Kilogram(p / 2.204)
}

func KilogramtoP(k Kilogram) Pound {
	return Pound(k * 2.204)
}
