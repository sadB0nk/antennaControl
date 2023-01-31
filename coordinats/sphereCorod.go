package coordinats

import (
	"fmt"
	"sync"
)

type sphereCoord struct {
	az float64
	el float64
	mu sync.Mutex
}

func (c sphereCoord) String() string { // Нужно для автоматического форматирования данных в консоль fmt.Println(sphereCoord)

	return fmt.Sprintf("Az:%g	El:%g", c.az, c.el)
}

func (c sphereCoord) Get() string {
	return fmt.Sprintf("%.2f\n%.4f\n", c.az, c.el) // Az
}

func (c sphereCoord) GetMassive() (float64, float64) {
	return c.az, c.el
}

func (c *sphereCoord) Set(az, el float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.az = az
	c.el = el
}
