package massconv

// KToP converts kilograms to pounds.
func KToP(k Kilogram) Pound { return Pound(k * Multiplier) }

// PToK converts pounds to kilograms
func PToK(p Pound) Kilogram { return Kilogram(p / Multiplier) }
