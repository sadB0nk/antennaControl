package main

import (
	"fmt"
	"sync"
)

// структура координат и методы доступа
type Coordinats struct {
	az float64
	el float64
	x  float64
	y  float64
	mu sync.Mutex
}

// дальше должна добавится блокировка для многопотока
func (c Coordinats) String() string { // Нужно для автоматического форматирования данных в консоль fmt.Println(Coordinats)
	c.mu.Lock()
	defer c.mu.Unlock()
	return fmt.Sprintf("Az:%g	El:%g", c.az, c.el)
}
func (c Coordinats) GetSpherePos() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return fmt.Sprintf("%g\n%g\n", c.az, c.el) // Az\nEl\n
}
func (c Coordinats) GetSpherePositionMassive() []float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return []float64{c.az, c.el}
}

func (c *Coordinats) SetPositon(az, el float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.az = az
	c.el = el
}
