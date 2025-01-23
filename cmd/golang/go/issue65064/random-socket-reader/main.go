// usage: random-socket-reader no._of_sockets iterations bytes
package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	logEnabled := flag.Bool("log", false, "enable logging")
	flag.Parse()

	if !*logEnabled {
		log.SetOutput(io.Discard)
	}

	if n := flag.NArg(); n != 3 {
		log.Fatalf("not enough/too many args (has %d, needs 3)", n)
	}

	nSockets, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Panic(err)
	}
	iterations, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Panic(err)
	}
	nBytes, err := strconv.ParseInt(flag.Arg(2), 10, 64)
	if err != nil {
		log.Panic(err)
	}

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		log.Panic(err)
	}
	defer os.RemoveAll(dir)

	randBuf := bytes.NewBuffer(nil)
	if _, err := io.CopyN(randBuf, rand.Reader, nBytes); err != nil {
		log.Panic(err)
	}

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var sockets []string
	for i := 0; i < nSockets; i++ {
		sockets = append(sockets, filepath.Join(dir, strconv.Itoa(i)))
	}

	ready := &sync.WaitGroup{}
	ready.Add(nSockets)
	done := make(chan struct{})

	var servers errgroup.Group
	for _, socket := range sockets {
		socket := socket
		servers.Go(func() error {
			startServer(socket, ready, done)
			return nil
		})
	}

	ready.Wait()

	var clients errgroup.Group
	for _, socket := range sockets {
		socket := socket
		clients.Go(func() error {
			startClient(socket, iterations, randBuf.Bytes())
			return nil
		})
	}

	clients.Wait()
	close(done)
	servers.Wait()
}

func startServer(socket string, ready *sync.WaitGroup, done <-chan struct{}) {
	listener, err := net.Listen("unix", socket)
	if err != nil {
		log.Panicf("server (%s): failed to listen on socket: %v", socket, err)
	}

	var g errgroup.Group

	g.Go(func() error {
		<-done
		return listener.Close()
	})

	log.Printf("server (%s): listening for connections…", socket)
	ready.Done()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server (%s): failed to accept connection: %v", socket, err)
			break
		}
		g.Go(func() error {
			handleConnection(socket, conn)
			return nil
		})
	}

	log.Printf("server (%s): closed: %v", socket, g.Wait())
}

func handleConnection(socket string, conn net.Conn) {
	defer conn.Close()
	log.Printf("server (%s): connection accepted", socket)

	n, err := io.Copy(io.Discard, conn)
	if err != nil {
		log.Printf("server (%s): connection closed?: %v", socket, err)
	}
	log.Printf("server (%s): received %d bytes", socket, n)
}

func startClient(socket string, iterations int, buf []byte) {
	for i := 0; i < iterations; i++ {
		sendData(socket, i, buf)
	}
}

func sendData(socket string, i int, buf []byte) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Printf("client (%s/%d): failed to connect to server: %v", socket, i, err)
	}
	defer conn.Close()
	log.Printf("client (%s/%d): connected to server", socket, i)

	if _, err = io.Copy(conn, bytes.NewBuffer(buf)); err != nil {
		log.Printf("client (%s): failed to write data: %v", socket, err)
	}
	log.Printf("client (%s/%d): random data sent", socket, i)
}
