package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/SwirlGit/md5-request-tool/hash"
	"github.com/SwirlGit/md5-request-tool/httpclient"
)

const defaultParallel = 10

func main() {
	var parallel int
	flag.IntVar(&parallel, "parallel", defaultParallel, "amount of parallel requests")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return
	}

	httpClient := httpclient.NewClient(httpclient.Config{
		Timeout:             5 * time.Second,
		IdleConnTimeout:     300 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	})
	hashService := hash.NewService(httpClient)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, parallel)
	for i := range args {
		i := i
		semaphore <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				<-semaphore
			}()

			url := addHTTPPrefix(args[i])
			md5Hash, err := hashService.MD5(context.Background(), url)
			if err != nil {
				fmt.Printf("%s %s\n", url, err)
				return
			}
			fmt.Printf("%s %x\n", url, md5Hash)
		}()
	}
	wg.Wait()
}

func addHTTPPrefix(url string) string {
	const prefix = "http://"
	if strings.HasPrefix(url, prefix) {
		return url
	}
	return prefix + url
}
