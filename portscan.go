package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var ipToScan string
var minPort = 1
var maxPort = 20000

func main() {
	fmt.Print("Введите ip адрес: ")
	fmt.Scanf("%s", &ipToScan)
	
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <=  maxPort; port++ {
		go testTCPConnection(ipToScan, port, doneChannel) // <-- go threads
		activeThreads++
	}

	// Ждем пока все ядра завершатся
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Host %s has open port: %d\n", ip, port)
	}
	
	doneChannel <- true
}









