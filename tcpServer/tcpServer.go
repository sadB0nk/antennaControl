package tcpServer

import (
	"antennaControl/coordinats"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

type tcpServ struct {
	d data
}
type Server interface {
	Start(coord *coordinats.Coord) (err error)
}

func Start(coord *coordinats.Coord) (err error) {
	var s tcpServ
	return s.start(coord)
}
func (s tcpServ) start(coord *coordinats.Coord) (err error) {

	const (
		getPos   byte = 'p'
		setPos   byte = 'P'
		shutdown byte = 'S'
	)
	log.Printf("Starting tcp server")
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		return err
	}
	defer listener.Close()
	log.Printf("Server started")
	data := s.d
	data.Bytes = make([]byte, 256)
	data.Strings = make([]string, 256)

	for { // Дальше это должно стать отдельным потоком
		log.Printf("Starting accepting")
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		log.Printf("Connect to ip = %v", conn.RemoteAddr())

		for {
			err = data.get(conn) // получение данных
			if err != nil {
				conn.Close()
				break
			}
			// обработка данных взасимости от данных приложения
			if data.Strings[0][0] == getPos {
				_, err := io.WriteString(conn, coord.RectToSphere().Get())
				if err != nil {
					return err
				}
			}
			// закрытые соединения
			if data.Strings[0][0] == shutdown {
				log.Println("Было разорвано соединение с Gpredict")
				_, err = io.WriteString(conn, "Ok")
				if err != nil {
					return err
				}
				conn.Close()
				break
			}
			// сет координат
			if data.Strings[0][0] == setPos {
				az, err := strconv.ParseFloat(data.Strings[1], 64) // конвертирование  строки в float
				if err != nil {
					log.Println(err)
					continue
				}
				el, err := strconv.ParseFloat(fmt.Sprintf(data.Strings[2]), 64)
				if err != nil {
					log.Println(err)
					continue
				}
				_, err = io.WriteString(conn, "Ok")
				if err != nil {
					log.Println(err)
					continue
				}
				coord.S.Set(az, el)
			}
		}
	}
}
