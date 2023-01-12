package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Data struct {
	Data []byte
	S    []string
}

func (d *Data) DataF() {
	tmp := strings.Split(string(d.Data), " ")
	for i, string2 := range tmp {
		s := ""
		for i := 0; i < len(string2); i++ {
			if string2[i] != '\x00' {
				if string2[i] != 44 {
					s += string(string2[i])
				} else {
					s += string(46)
				}
			}
		}
		d.S[i] = s
	}
	return
}

func (d *Data) Get(conn net.Conn) {
	_, err := conn.Read(d.Data)
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}
	log.Println(d.Data)
	d.DataF()
	log.Println(d.S)
	return
}

func main() {
	const (
		getPos   byte = 'p'
		setPos   byte = 'P'
		shutdown byte = 'S'
	)

	data := Data{make([]byte, 256),
		make([]string, 4)}
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
			data.Get(conn)
			if data.S[0][0] == getPos {
				_, err := io.WriteString(conn, coord.GetPosition())
				if err != nil {
					return
				}
			}

			if data.S[0][0] == shutdown {
				fmt.Println("Было разорвано соединение с Gpredict")
				conn.Close()
				break
			}

			if data.S[0][0] == setPos {
				az, err := strconv.ParseFloat(data.S[1], 64)
				if err != nil {
					log.Println(err)
					continue
				}
				el, err := strconv.ParseFloat(fmt.Sprintf(data.S[2]), 64)
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
