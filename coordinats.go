package main

import (
	"fmt"
)

type Coordinats struct {
	Az float64
	El float64
}

func (c Coordinats) String() string {
	return fmt.Sprintf("Az:%g	El:%g", c.Az, c.El)
}
func (c Coordinats) GetPosition() string {
	return fmt.Sprintf("%g\n%g\n", c.Az, c.El) // Az\nEl\n
}
func (c Coordinats) GetPositionMassive() []float64 {
	return []float64{c.Az, c.El}
}

func (c *Coordinats) SetPositon(az, el float64) {
	c.Az = az
	c.El = el
}
