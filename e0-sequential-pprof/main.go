package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"
)

func main() {
	pf, err := os.Create("e0.goroutine.prof")
	if err != nil {
		log.Fatalln(err)
	}
	defer pf.Close()
	defer pprof.Lookup("goroutine").WriteTo(pf, 0)

	for i := 5300; i <= 5500; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", i))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", i, err)
			continue
		}
		conn.Close()
		log.Printf("%d OPEN\n", i)
	}
}
