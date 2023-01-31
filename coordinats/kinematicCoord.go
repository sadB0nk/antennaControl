package coordinats

import (
	"fmt"
	"sync"
)

type kinematicCoord struct {
	firstAxis  float64
	secondAxis float64
	mu         sync.Mutex
}

func (r kinematicCoord) String() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return fmt.Sprintf("X:%.4f	Y:%.4f", r.firstAxis, r.secondAxis)
}

func (r kinematicCoord) Get() string {
	return fmt.Sprintf("%.4f\n%.4f\n", r.firstAxis, r.secondAxis) // Az
}

func (r kinematicCoord) GetMassive() (float64, float64) {
	return r.firstAxis, r.secondAxis
}

func (r *kinematicCoord) Set(firstax, secondax float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.firstAxis = firstax
	r.secondAxis = secondax
}
