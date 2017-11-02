package conv

import "fmt"

type Pond float64
type Kilo float64

func (p Pond) String() string { return fmt.Sprintf("%g lb", p) }
func (k Kilo) String() string { return fmt.Sprintf("%g kg", k) }

func PToK(p Pond) Kilo { return Kilo(p / 2.2046) }
func KToP(k Kilo) Pond { return Pond(k * 2.2046) }
