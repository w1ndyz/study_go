package tempconv

import "fmt"

type Kelvin float64
type Celsius float64

func (k Kelvin) String() string {
	return fmt.Sprintf("%gk", k)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}
