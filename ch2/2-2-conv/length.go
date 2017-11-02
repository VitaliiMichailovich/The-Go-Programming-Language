package conv

import "fmt"

type Foot float64
type Metr float64

func (f Foot) String() string { return fmt.Sprintf("%g ft", f) }
func (m Metr) String() string { return fmt.Sprintf("%g m", m) }

func FToM(f Foot) Metr { return Metr(f / 3.2808) }
func MToF(m Metr) Foot { return Foot(m * 3.2808) }
