package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	fileReader()
	stringReader()
	connReader()
}

func fileReader() {

	f, err := os.Open("read.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buffer := make([]byte, 10)

	for {

		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}
	}
}

func stringReader() {

	s := strings.NewReader("this is a test string")

	buffer := make([]byte, 1)

	for {

		n, err := s.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}

	}

}

func connReader() {

	conn, err := net.Dial("tcp", "google.com:80")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET HTTP 1.0\r\n\r\n")

	buffer := make([]byte, 25)

	for {

		n, err := conn.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}

	}
}

func readerToStdout(r io.Reader, buffSize int) {
	buffer := make([]byte, buffSize)

	for {

		n, err := r.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		if n > 0 {
			fmt.Println(string(buffer[:n]))
		}
	}
}
