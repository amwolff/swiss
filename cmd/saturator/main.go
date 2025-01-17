// usage: saturator list.txt repeat
package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {
	flag.Parse()

	if n := flag.NArg(); n != 2 {
		log.Fatalf("not enough/too many args (has %d, needs 2)", n)
	}

	repeat, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Panic(err)
	}

	data, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		log.Panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var group errgroup.Group
	for _, u := range lines {
		for i := 0; i < repeat; i++ {
			i, u := i, u
			group.Go(func() error {
				resp, err := http.Get(u)
				if err != nil {
					return err
				}
				log.Printf("[%d] starting streaming from %s", i, u)
				n, err := io.Copy(io.Discard, resp.Body)
				log.Printf("[%d] downloaded %d bytes from %s (err: %v)", i, n, u, err)
				return errors.Join(err, resp.Body.Close())
			})
		}
	}
	if err = group.Wait(); err != nil {
		log.Printf("finished with error: %v", err)
	} else {
		log.Println("finished without errors")
	}
}
