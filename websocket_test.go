package okx

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestWebsocket(t *testing.T) {
	client := NewClient(API_KEY, SECRET_KEY, PASSPHRASE)

	ch := make(chan *WsOrderDetail)
	err := client.SubscribeOrderChannel(ch)
	require.NoError(t, err)

	go func(ch chan *WsOrderDetail) {
		for {
			select {
			case event := <-ch:
				fmt.Println("recv", event)
			}
		}
	}(ch)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
