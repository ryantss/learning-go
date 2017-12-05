package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("filename", "", "File name")
	flag.Parse()

	buf := make([]byte, 1024)
	f, e := os.Open(*filename)
	if e != nil {
		log.Fatalf(e.Error())
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		n, e := r.Read(buf)
		if e != nil {
			log.Fatalf(e.Error())
		}
		if n == 0 {
			break
		}
		fmt.Printf("%s\n", string(buf))
	}

}
