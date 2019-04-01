package lenconv

// MToF converts meters to feets.
func MToF(m Meter) Feet { return Feet(m * 3.28083) }

// FToM converts feets to meters.
func FToM(f Feet) Meter { return Meter(f / 3.28083) }
