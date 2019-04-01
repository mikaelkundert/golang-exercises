package lenconv

import "fmt"

// Meter is one of the distance types.
type Meter float64

// Feet is one of the distance types.
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
