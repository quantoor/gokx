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

	eventHandler := func(event *WsOrdersEvent) {
		for _, v := range event.Data {
			fmt.Println(v)
		}
	}

	// if a websocket has an error
	errHandler := func(err error) {
		fmt.Println(err)
	}
	err := client.SubscribeOrderEvents(eventHandler, errHandler)
	require.NoError(t, err)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
