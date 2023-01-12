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
	Bytes   []byte
	Strings []string
}

// Форматирование данных, заменг запятых на точки и уборка нулевых символов
func (d *Data) DataF() {
	tmp := strings.Split(string(d.Bytes), " ")
	for i, string2 := range tmp {
		s := ""
		for i := 0; i < len(string2); i++ {
			if string2[i] != '\x00' {
				if string2[i] != 44 {
					s += string(string2[i])
				} else {
					s += string('.')
				}
			}
		}
		d.Strings[i] = s
	}
	return
}

// Получение данных от Gpredict
func (d *Data) Get(conn net.Conn) (err error) {
	tmp := make([]byte, 256)
	_, err = conn.Read(tmp)
	d.Bytes = tmp
	if err != nil {
		log.Println(err)
		return err
	}
	d.DataF()
	return nil
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
	var coord Coordinats // инициализация структуры координат
	coord.Az = 17
	coord.El = 12
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
			if data.Strings[0][0] == getPos { // обработка данных взасимости от данных приложения
				_, err := io.WriteString(conn, coord.GetPosition())
				if err != nil {
					return
				}
			}

			if data.Strings[0][0] == shutdown {
				log.Println("Было разорвано соединение с Gpredict") // закрытые соединения
				_, err = io.WriteString(conn, "Ok")
				if err != nil {
					log.Println(err)
					log.Fatal()
				}
				conn.Close()
				break
			}
			if data.Strings[0][0] == setPos { // сет координат
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
