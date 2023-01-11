package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
func getData(conn net.Conn) []string {
	data := make([]byte, 256)
	_, err := conn.Read(data)
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}
	return strings.Split(string(data), " ")
}
func main() {
	const (
		getPos   string = "p"
		setPos   string = "P"
		shutdown string = "S"
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
			fmt.Printf("%v\n", data[0])
			fmt.Printf("%v\n", data[0] == "p")
			fmt.Printf("%s\n", data[0] == getPos)
			if data[0] == getPos {
				_, err := io.WriteString(conn, coord.GetPosition())
				if err != nil {
					return
				}
			}

			if data[0] == shutdown {
				fmt.Println("Было разорвано соединение с Gpredict")
				conn.Close()
				break
			}
		}
	}
}
