package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/marcy-t/Sample-Websocket-from-golang/infra/websocket"
	"github.com/marcy-t/Sample-Websocket-from-golang/intefaces"
)

const (
	defaultGOMAXPROCS int = 2
)

var (
	cpu = flag.Int("c", defaultGOMAXPROCS, "number of GOMAXPROCS")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*cpu)
	fmt.Println("webcoket sample")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	// キャンセル待ちうけるためのgoroutine
	go func() {
		defer close(sigChan)
		for sig := range sigChan {
			fmt.Printf("SIGNAL %d received. then canceling jobs ...", sig)
			cancel()
			break
		}
	}()

	ws, err := websocket.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// キャンセルされるまで処理をしつづけるgoroutine
	ne := intefaces.New(ws)
	ne.Run(ctx)
}
