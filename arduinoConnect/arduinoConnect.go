package arduinoConnect

import (
	"antennaControl/coordinats"
	"bufio"
	"go.bug.st/serial"
	"io"
	"log"
	"strconv"
)

func dataF(string2 string) (s string) {

	for i := 0; i < len(string2); i++ {
		if string2[i] != '\x00' && string2[i] != '\n' && string2[i] != '\r' {
			if string2[i] != 44 {
				s += string(string2[i])
			} else {
				s += string('.')
			}
		}
	}
	return
}

type Server interface {
	Start(coord *coordinats.Coord) (err error)
}
type Arduino struct {
}

var a Arduino

func Start(coord *coordinats.Coord) (err error) {

	return a.start(coord)
}
func (Arduino) start(coord *coordinats.Coord) (err error) {
	log.Printf("Starting Arduino connect")
	mode := &serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	}

	conn, err := serial.Open("COM3", mode)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Arduino wac connected")
	for {
		io.WriteString(conn, coord.SphereToRect().Get())
		buff := bufio.NewReader(conn)
		tmp, err := buff.ReadString('\n')

		if err != nil {
			return err
		}
		x, err := strconv.ParseFloat(dataF(tmp), 64)
		if err != nil {
			continue
		}
		tmp, err = buff.ReadString('\n')
		y, err := strconv.ParseFloat(dataF(tmp), 64)
		if err != nil {
			continue
		}
		coord.R.Set(x, y)
	}
}
