package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"
	"strconv"
	"sync"
	"time"
)

var host string
var fromPort string
var toPort string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "Host to scan.")
	flag.StringVar(&fromPort, "from", "5300", "Port to start scanning from")
	flag.StringVar(&toPort, "to", "5500", "Port at which to stop scanning")
}

func main() {
	flag.Parse()

	fp, err := strconv.Atoi(fromPort)
	if err != nil {
		log.Fatalln("Invalid 'from' port")
	}

	tp, err := strconv.Atoi(toPort)
	if err != nil {
		log.Fatalln("Invalid 'to' port")
	}

	if fp > tp {
		log.Fatalln("Invalid values for 'from' and 'to' port")
	}

	pf, err := os.Create("e1.goroutine.prof")
	if err != nil {
		log.Fatalln(err)
	}
	defer pf.Close()

	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		time.Sleep(250 * time.Millisecond)
		if err := pprof.Lookup("goroutine").WriteTo(pf, 0); err != nil {
			log.Println(err)
		}
	}()

	wg.Add(tp - fp + 1) // +1 to avoid an off by one issue
	wg.Add(1)           // +1 for the goroutine that writes to the file

	for i := fp; i <= tp; i++ {
		go func(p int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, p))
			if err != nil {
				log.Printf("%d CLOSED (%s)\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d OPEN\n", p)
		}(i)
	}
	wg.Wait()
}