package main

import (
	"antennaControl/coordinats"
	"fmt"
	"time"
)

func debug(coord *coordinats.Coord) {

	for {
		time.Sleep(time.Millisecond * 300)

		az, el := coord.S.GetMassive()
		x, y := coord.R.GetMassive()
		fmt.Printf("Az: %.4f, El: %.4f 	X: %.4f, Y: %.4f", az, el, x, y)
		az, el = coord.RectToSphere().GetMassive()
		x, y = coord.SphereToRect().GetMassive()
		f, s := coord.RectToKinematic().GetMassive()

		fmt.Printf("	RectfromSpher: Az: %.4f, El: %.4f  SphereFromRect: X: %.4f, Y: %.4f  RectToKinemati: F: %.4f, S: %.4f\n", az, el, x, y, f, s)

	}

}
