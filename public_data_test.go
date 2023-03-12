package okx

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetInstruments(t *testing.T) {
	client := NewClient("", "", "")
	_, err := client.NewGetInstrumentsService().
		InstrumentType("SWAP").
		Do(context.Background())
	require.NoError(t, err)
}

func TestGetLimitPrice(t *testing.T) {
	client := NewClient("", "", "")
	res, err := client.NewGetLimitPriceService().
		InstrumentId("MATIC-USDT-SWAP").
		Do(context.Background())
	fmt.Println(res)
	require.NoError(t, err)
}
