package main

import (
	"os"
	"fmt"
	"log"
	"net"
	"time"
	"flag"
	"strconv"
	"moul.io/banner"
    "github.com/fatih/color"
)

func banner_start() {
	fmt.Println(banner.Inline("portssscaner"))
	fmt.Println("\t\t\t\t\tAuthor: quietmoth1")
}

func main() {
	banner_start()
	var minPort = 1          // minport value 

	// ARGS
	ipToScan := flag.String("ip", "", "IP_to_scan")
	maxPort := flag.CommandLine.Int("ports", 0, "number of ports to scan")

	flag.Parse()

	// Check args
	if len(*ipToScan) == 0 || *maxPort < 1 || *maxPort > 65535 {
		fmt.Println("Usage: portssscanner -ip 8.8.8.8")
		flag.PrintDefaults()
		os.Exit(0)
	}
	 

	fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// Threads
	activeThreads := 0
	doneChannel := make(chan bool)

	for port := minPort; port <= *maxPort; port++ {
		go TCPConnectionScan(*ipToScan, port, doneChannel) // <-- go threads
		activeThreads++
	}

	// Wait to threads are stop
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
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
