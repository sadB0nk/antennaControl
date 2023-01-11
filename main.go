package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

/*
	func readBytes(reader *bufio.Reader) string {
		for {
			bytik, err := reader.ReadByte()
			if err != nil {
				break
			}
			data = append(data, bytik)
			if bytik == 112 {
				break
			}
		}
		return "tmp"
	}
*/
func main() {
	const (
		getPos   byte = 112
		setPos   byte = 80
		shutdown byte = 83
		buffSize      = 256
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
	fmt.Println(coord.GetPositionMassive())
	data := make([]byte, 256)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Connect to ip = %v", conn.RemoteAddr())
		//	buff := bufio.NewReader(conn)
		for {
			fmt.Println(coord.GetPositionMassive())
			_, err = conn.Read(data)
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Println(data)
			if data[0] == getPos {
				_, err := io.WriteString(conn, coord.GetPosition())
				if err != nil {
					return
				}
			}
			log.Printf("Recieved data: %s", data)

			log.Printf("Recieved data: %v", data)
			if data[0] == 83 {
				fmt.Println("Было разорвано соединение с Gpredict")
				conn.Close()
				break
			}

		}
	}
}
