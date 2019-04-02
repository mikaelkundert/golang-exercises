package massconv

import "fmt"

// Kilogram is a type of weight unit.
type Kilogram float64

// Pound is a type of weight unit.
type Pound float64

// Multiplier specifies the ratio between Kilograms and pounds
const Multiplier = 2.204623

func (k Kilogram) String() string { return fmt.Sprintf("%gKg", k) }
func (p Pound) String() string    { return fmt.Sprintf("%glbs", p) }
