package lenconv

// MToF converts meters to feets.
func MToF(m Meter) Feet { return Feet(m * Multiplier) }

// FToM converts feets to meters.
func FToM(f Feet) Meter { return Meter(f / Multiplier) }
