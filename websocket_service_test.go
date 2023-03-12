package okx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestWebsocket(t *testing.T) {
	// when a Balance Position message of the websocket has arrived is invoked
	wsHandlerBalancePosition := func(event *WsBalancePositionEvent) {
		for _, v := range event.Data {
			fmt.Println(v)
		}
	}

	// when a Balance Position message of the websocket has arrived is invoked
	wsHandlerOrders := func(event *WsOrdersEvent) {
		reqBodyBytesTemp := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytesTemp).Encode(&event)
		for _, v := range event.Data {
			fmt.Println(v)
		}
	}

	// if a websocket has an error
	errHandlerBalancePosition := func(err error) {
		fmt.Println(err)
	}

	// if a websocket has an error
	errHandlerOrders := func(err error) {
		fmt.Println(err)
	}

	// Launch coroutine to listen BalancePosition updates
	go func() {
		_, _, err := WsBalancePositionServe(API_KEY, SECRET_KEY, PASSPHRASE, wsHandlerBalancePosition, errHandlerBalancePosition, false)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Launch coroutine to listen Orders updates
	go func() {
		_, _, err := WsOrdersServe("ANY", "", "", API_KEY, SECRET_KEY, PASSPHRASE, wsHandlerOrders, errHandlerOrders, false)
		if err != nil {
			fmt.Println(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
