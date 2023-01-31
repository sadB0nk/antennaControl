package coordinats

import (
	"math"
)

// структура координат и методы доступа

type Coord struct {
	S sphereCoord
	R rectanCoord
	K kinematicCoord
}

func (Coord) rad(degr float64) float64 {
	return degr * (math.Pi / 180)
}

func (c Coord) elev() float64 {
	return math.Sqrt(c.R.x*c.R.x + c.R.y*c.R.y)
}

func (Coord) degr(rad float64) float64 {
	return rad / (math.Pi / 180)
}

func (c Coord) RectToSphere() sphereCoord {
	c.S.az = math.Asin(c.R.x/c.elev()) / (math.Pi / 180)
	c.S.el = math.Acos(c.elev()) / (math.Pi / 180)
	if c.R.x > 0 && c.R.y < 0 {

		c.S.az = 360 - c.S.az - 180
	}
	if c.R.x <= 0 && c.R.y < 0 {
		c.S.az = 360 - c.S.az + 180 - 360
	}
	if c.R.x < 0 && c.R.y > 0 {
		c.S.az += 360
	}
	return c.S
}

func (c Coord) SphereToRect() rectanCoord {
	c.R.y = math.Cos(c.rad(c.S.az)) * math.Cos(c.rad(c.S.el))
	c.R.x = math.Sin(c.rad(c.S.az)) * math.Cos(c.rad(c.S.el))
	return c.R
}

func (c Coord) RectToKinematic() (k kinematicCoord) {
	gyp := math.Sqrt(c.R.x*c.R.x + c.R.y*c.R.y)
	q1 := math.Atan2(c.R.y, c.R.x)
	q2 := math.Acos((0.5 - gyp*gyp) / gyp)
	k.firstAxis = c.degr(q1 + q2)
	k.secondAxis = c.degr((0.5 - gyp*gyp) / 0.5)
	return k
}
