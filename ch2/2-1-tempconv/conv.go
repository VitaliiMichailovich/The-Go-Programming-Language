package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) / 1.8) }

// CToF converts a Celsius temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((9/5)*(k-273.15) + 32) }

// CToF converts a Celsius temperature to Fahrenheit.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
