package main

import (
	"antennaControl/arduinoConnect"
	"antennaControl/coordinats"
	"antennaControl/tcpServer"
	"log"
	"sync"
)

func main() {
	var sw sync.WaitGroup
	var coord coordinats.Coord
	sw.Add(1)
	go func() {
		defer sw.Done()
		err := tcpServer.Start(&coord)
		if err != nil {
			log.Fatal(err)
		}
	}()
	sw.Add(1)
	go func() {
		defer sw.Done()
		err := arduinoConnect.Start(&coord)
		if err != nil {
			log.Fatal(err)
		}
	}()
	sw.Add(1)
	go debug(&coord)
	sw.Wait()
}
