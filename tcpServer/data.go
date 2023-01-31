package tcpServer

import (
	"io"
	"log"
	"strings"
)

type data struct {
	Bytes   []byte
	Strings []string
}

func (d *data) dataF() {
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
func (d *data) get(conn io.Reader) (err error) {
	tmp := make([]byte, 256)
	_, err = conn.Read(tmp)
	d.Bytes = tmp
	if err != nil {
		log.Println(err)
		return err
	}
	d.dataF()
	return nil
}
