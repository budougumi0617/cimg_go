package cimg_go

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	var srv http.Server
	var mu sync.Mutex
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("start listen", l.Addr().String())
	oldHook := srv.ConnState
	srv.ConnState = func(c net.Conn, cs http.ConnState) {
		mu.Lock()
		defer mu.Unlock()

		fmt.Println("conn status", cs.String())

		if oldHook != nil {
			oldHook(c, cs)
		}
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		res, err := http.Get("http://" + l.Addr().String())
		if err != nil {
			log.Fatal(err)
		}
		robots, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)
		time.Sleep(1 * time.Second)

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.Serve(l); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
