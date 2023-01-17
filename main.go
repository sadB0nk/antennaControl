package main

import "sync"

var sw sync.WaitGroup
var coord Coordinats
var server tcpServ

func main() {

	sw.Add(1)
	go server.Start()
	sw.Wait()

}
