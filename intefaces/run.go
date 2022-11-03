package intefaces

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcy-t/Sample-Websocket-from-golang/infra/websocket"
)

type NewHandler struct {
	WebSocket websocket.Interface
}

type Handler interface {
	Run(ctx context.Context)
}

func New(ws websocket.Interface) Handler {
	return &NewHandler{
		WebSocket: ws,
	}
}

func (h *NewHandler) Run(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer close(sigChan)
		for sig := range sigChan {
			fmt.Sprintf("SIGNAL %d received. then canceling jobs !!!!! ...", sig)
			break
		}
	}()

	ticker := time.NewTicker(2 * time.Second)
	func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Sprintf("Shutting down.!!!!!!!!")
				return
			case <-ticker.C:
				h.receive(ctx)
			}
		}
	}()
}

func (h *NewHandler) receive(ctx context.Context) (err error) {
	fmt.Println("##################################")
	fmt.Println("START")
	h.WebSocket.Con()
	fmt.Println("##################################")
	return
}
