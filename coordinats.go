package main

import (
	"fmt"
)

// структура координат и методы доступа
type Coordinats struct {
	Az float64
	El float64
}

// дальше должна добавится блокировка для многопотока
func (c Coordinats) String() string { // Нужно для автоматического форматирования данных в консоль fmt.Println(Coordinats)
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
