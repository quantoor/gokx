package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetTickerService
type GetTickerService struct {
	c      *Client
	instId string
}

func (s *GetTickerService) InstrumentId(instId string) *GetTickerService {
	s.instId = instId
	return s
}

func (s *GetTickerService) Do(ctx context.Context, opts ...RequestOption) (res *GetTickerServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/market/ticker",
	}

	r.setParam("instId", s.instId)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetTickerServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetTickerServiceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstType  string `json:"instType"`
		InstId    string `json:"instId"`
		Last      string `json:"last"`
		LastSz    string `json:"lastSz"`
		AskPx     string `json:"askPx"`
		AskSz     string `json:"askSz"`
		BidPx     string `json:"bidPx"`
		BidSz     string `json:"bidSz"`
		Open24H   string `json:"open24h"`
		High24H   string `json:"high24h"`
		Low24H    string `json:"low24h"`
		VolCcy24H string `json:"volCcy24h"`
		Vol24H    string `json:"vol24h"`
		SodUtc0   string `json:"sodUtc0"`
		SodUtc8   string `json:"sodUtc8"`
		Ts        string `json:"ts"`
	} `json:"data"`
}
