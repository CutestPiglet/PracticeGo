package lengthconv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g Feet", f) }
func (m Meters) String() string { return fmt.Sprintf("%g Meters", m) }
