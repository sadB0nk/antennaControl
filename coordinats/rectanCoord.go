package coordinats

import (
	"fmt"
	"sync"
)

type rectanCoord struct {
	x  float64
	y  float64
	mu sync.Mutex
}

func (r rectanCoord) String() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return fmt.Sprintf("X:%.4f	Y:%.4f", r.x, r.y)
}

func (r rectanCoord) Get() string {
	return fmt.Sprintf("%.4f\n%.4f\n", r.x, r.y) // Az
}

func (r rectanCoord) GetMassive() (float64, float64) {
	return r.x, r.y
}

func (r *rectanCoord) Set(x, y float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.x = x
	r.y = y
}
