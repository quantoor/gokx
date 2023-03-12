package okx

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTicker(t *testing.T) {
	client := NewClient("", "", "")
	res, err := client.NewGetTickerService().
		InstrumentId("MATIC-USDT-SWAP").
		Do(context.Background())
	require.NoError(t, err)
	fmt.Println(res)
}
