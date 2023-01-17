package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

type tcpServ struct {
}

func (tcpServ) Start() {
	defer sw.Done()
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
	data := Data{make([]byte, 256),
		make([]string, 4)}

	for { // Дальше это должно стать отдельным потоком
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connect to ip = %v", conn.RemoteAddr())

		for {
			err = data.Get(conn) // получение данных
			if err != nil {
				conn.Close()
				break
			}
			// обработка данных взасимости от данных приложения
			if data.Strings[0][0] == getPos {
				_, err := io.WriteString(conn, coord.GetSpherePos())
				if err != nil {
					return
				}
			}
			// закрытые соединения
			if data.Strings[0][0] == shutdown {
				log.Println("Было разорвано соединение с Gpredict")
				_, err = io.WriteString(conn, "Ok")
				if err != nil {
					log.Println(err)
					log.Fatal()
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
				coord.SetPositon(az, el)
				log.Println(coord)
			}
		}
	}
}
