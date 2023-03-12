package okx

import (
	"context"
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
	_, err := client.NewGetLimitPriceService().
		InstrumentId("SWAP").
		Do(context.Background())
	require.NoError(t, err)
}
