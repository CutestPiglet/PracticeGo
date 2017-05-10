package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%g Pound", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g Kilogram", k) }
