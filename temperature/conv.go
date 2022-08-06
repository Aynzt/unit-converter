package temperature

//  CtoF convert Celsuis to Fahrenheit
func CtoF(c Celsuis) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

//  FtoC convert Fahrenheit to Celsuis
func FtoC(f Fahrenheit) Celsuis {
	return Celsuis(f-32) * 5 / 9
}

// CToK convert Celsius to Kelvin
func CtoK(c Celsuis) Kelvin {
	return Kelvin(c + 273)
}

// KtoC convert kelvin to celsius
func KtoC(k Kelvin) Celsuis {
	return Celsuis(k - 273)
}
