package okx

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var client = NewClient(API_KEY, SECRET_KEY, PASSPHRASE)

const (
	instrumentId = "MATIC-USDT-SWAP"
)

func TestPlaceOrder(t *testing.T) {
	res, err := client.NewPlaceOrderService().
		InstrumentId(instrumentId).
		TradeMode("cross").
		Side("buy").
		OrderType("limit").
		Size("1").
		Price("0.95").
		Do(context.Background())
	require.NoError(t, err)

	orderId := res.Data[0].OrdId
	_, err = client.NewCancelOrderService().
		InstrumentId(instrumentId).
		OrderId(orderId).
		Do(context.Background())
	require.NoError(t, err)
}

func TestAmendOrder(t *testing.T) {
	res, err := client.NewPlaceOrderService().
		InstrumentId(instrumentId).
		TradeMode("cross").
		Side("buy").
		OrderType("limit").
		Size("1").
		Price("0.95").
		Do(context.Background())
	require.NoError(t, err)

	orderId := res.Data[0].OrdId

	_, err = client.NewAmendOrderService().
		InstrumentId(instrumentId).
		OrderId(orderId).
		NewSize("2").
		NewPrice("0.9").
		Do(context.Background())
	require.NoError(t, err)

	_, err = client.NewCancelOrderService().
		InstrumentId(instrumentId).
		OrderId(orderId).
		Do(context.Background())
	require.NoError(t, err)
}

func TestCancelMultipleOrders(t *testing.T) {
	_, err := client.NewCancelMultipleOrdersService().
		InstrumentId(instrumentId).
		Do(context.Background())
	require.NoError(t, err)
}

func TestGetOrderList(t *testing.T) {
	res, err := client.NewGetOrderListService().
		InstrumentType(instrumentId).
		Do(context.Background())
	require.NoError(t, err)
	fmt.Println(res)
}
