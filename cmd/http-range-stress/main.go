// usage: http-range-stress URL Range_HTTP_request_header_value no._of_workers duration
package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
	"storj.io/common/sync2"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatalln("not enough args")
	}

	url, rangeValue := os.Args[1], os.Args[2]

	duration, err := time.ParseDuration(os.Args[4])
	if err != nil {
		log.Panicf("ParseDuration: %s", err)
	}
	workers, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Panicf("Atoi: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	var downloaded uint64

	group, groupCtx := errgroup.WithContext(ctx)
	for i := 0; i < workers; i++ {
		group.Go(func() error {
			for {
				select {
				case <-groupCtx.Done():
					return nil
				default:
					req, err := http.NewRequestWithContext(groupCtx, http.MethodGet, url, nil)
					if err != nil {
						return err
					}
					req.Header.Set("Range", rangeValue)

					// start := time.Now()
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						if errors.Is(err, context.DeadlineExceeded) {
							return nil
						}
						return err
					}
					defer resp.Body.Close()

					if resp.StatusCode != http.StatusPartialContent {
						log.Printf("unexpected response: %s", resp.Status)
						continue
					}

					n, err := sync2.Copy(groupCtx, io.Discard, resp.Body)
					if err != nil {
						return err
					}
					// log.Printf("downloaded %d in %s", n, time.Since(start))
					atomic.AddUint64(&downloaded, uint64(n))
				}
			}
		})
	}
	if err = group.Wait(); err != nil {
		log.Printf("finished with errors: %s", err)
	}
	log.Printf("downloaded %d bytes", downloaded)
}
