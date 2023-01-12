package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func dataF(string2 string) (s string) {
	for i := 0; i < len(string2); i++ {
		if string2[i] != '\x00' {
			if string2[i] != 44 {
				s += string(string2[i])
			} else {
				s += string(46)
			}
		}
	}
	return
}
func getData(conn net.Conn) []string {
	data := make([]byte, 256)
	_, err := conn.Read(data)
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}
	log.Println(data)
	return strings.Split(string(data), " ")
}

func main() {
	const (
		getPos   byte = 'p'
		setPos   byte = 'P'
		shutdown byte = 'S'
	)

	log.Printf("Starting tcp server")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	var coord Coordinats
	coord.Az = 17
	coord.El = 12
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connect to ip = %v", conn.RemoteAddr())

		for {
			data := getData(conn)
			if data[0][0] == getPos {
				_, err := io.WriteString(conn, coord.GetPosition())
				if err != nil {
					return
				}
			}

			if data[0][0] == shutdown {
				fmt.Println("Было разорвано соединение с Gpredict")
				conn.Close()
				break
			}

			if data[0][0] == setPos {
				az, err := strconv.ParseFloat(dataF(data[1]), 64)
				el, err := strconv.ParseFloat(dataF(data[2]), 64)
				if err != nil {
					log.Println(err)
					continue
				}
				coord.SetPositon(az, el)
				fmt.Println(coord)
			}
		}
	}
}
