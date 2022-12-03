package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
    "github.com/fatih/color"
	"moul.io/banner"
)

func banner_start() {
	fmt.Println(banner.Inline("portssscaner"))
	fmt.Println("\t\t\t\t\tAuthor: quietmoth1")
}

func check_args() (string, int, int) {
	fmt.Print("\nHello, give me your ip to scan ports: ")
	var ipToScan string
	fmt.Scanf("%s", &ipToScan)

	var minPort = 1          // minport value 
	var maxPort = 20000		// maxport value 
	return ipToScan, minPort, maxPort
}

func main() {
	banner_start()
	ipToScan, minPort, maxPort := check_args()

	
	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <=  maxPort; port++ {
		go testTCPConnection(ipToScan, port, doneChannel) // <-- go threads
		activeThreads++
	}

	// wait to thread are stop
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func TCPConnectionScan(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)

	ip = color.RedString(ip)
	color_port_string := strconv.Itoa(port)
	color_port_string = color.RedString(color_port_string)

	if err == nil {
		log.Printf("| Host %s has open port: %s \n", ip, color_port_string)
	} 
	
	doneChannel <- true
}
