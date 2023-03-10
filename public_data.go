package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetInstrumentsService struct {
	c        *Client
	instType string
}

func (s *GetInstrumentsService) InstrumentType(instType string) *GetInstrumentsService {
	s.instType = instType
	return s
}

func (s *GetInstrumentsService) Do(ctx context.Context, opts ...RequestOption) (res *GetInstrumentsServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/instruments",
	}

	r.setParam("instType", s.instType)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetInstrumentsServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetInstrumentsServiceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstType     string `json:"instType"`
		InstId       string `json:"instId"`
		InstFamily   string `json:"instFamily"`
		Uly          string `json:"uly"`
		Category     string `json:"category"`
		BaseCcy      string `json:"baseCcy"`
		QuoteCcy     string `json:"quoteCcy"`
		SettleCcy    string `json:"settleCcy"`
		CtVal        string `json:"ctVal"`
		CtMult       string `json:"ctMult"`
		CtValCcy     string `json:"ctValCcy"`
		OptType      string `json:"optType"`
		Stk          string `json:"stk"`
		ListTime     string `json:"listTime"`
		ExpTime      string `json:"expTime"`
		Lever        string `json:"lever"`
		TickSz       string `json:"tickSz"`
		LotSz        string `json:"lotSz"`
		MinSz        string `json:"minSz"`
		CtType       string `json:"ctType"`
		Alias        string `json:"alias"`
		State        string `json:"state"`
		MaxLmtSz     string `json:"maxLmtSz"`
		MaxMktSz     string `json:"maxMktSz"`
		MaxTwapSz    string `json:"maxTwapSz"`
		MaxIcebergSz string `json:"maxIcebergSz"`
		MaxTriggerSz string `json:"maxTriggerSz"`
		MaxStopSz    string `json:"maxStopSz"`
	} `json:"data"`
}
