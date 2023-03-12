package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

// PlaceOrderService
type PlaceOrderService struct {
	c       *Client
	instId  string
	tdMode  string
	side    string
	ordType string
	sz      string
	px      *string
}

func (s *PlaceOrderService) InstrumentId(instId string) *PlaceOrderService {
	s.instId = instId
	return s
}

func (s *PlaceOrderService) TradeMode(tdMode string) *PlaceOrderService {
	s.tdMode = tdMode
	return s
}

func (s *PlaceOrderService) Side(side string) *PlaceOrderService {
	s.side = side
	return s
}

func (s *PlaceOrderService) OrderType(ordType string) *PlaceOrderService {
	s.ordType = ordType
	return s
}

func (s *PlaceOrderService) Size(sz string) *PlaceOrderService {
	s.sz = sz
	return s
}

func (s *PlaceOrderService) Price(px string) *PlaceOrderService {
	s.px = &px
	return s
}

func (s *PlaceOrderService) Do(ctx context.Context, opts ...RequestOption) (res *PlaceOrderServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/order",
		secType:  secTypeSigned,
	}

	r.setBodyParam("instId", s.instId)
	r.setBodyParam("tdMode", s.tdMode)
	r.setBodyParam("side", s.side)
	r.setBodyParam("ordType", s.ordType)
	r.setBodyParam("sz", s.sz)

	if s.px != nil {
		r.setBodyParam("px", *s.px)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(PlaceOrderServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PlaceOrderServiceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ClOrdId string `json:"clOrdId"`
		OrdId   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}

// CancelOrderService
type CancelOrderService struct {
	c       *Client
	instId  string
	ordId   *string
	clOrdId *string
}

func (s *CancelOrderService) InstrumentId(instId string) *CancelOrderService {
	s.instId = instId
	return s
}

func (s *CancelOrderService) OrderId(ordId string) *CancelOrderService {
	s.ordId = &ordId
	return s
}

func (s *CancelOrderService) ClientOrderId(clOrdId string) *CancelOrderService {
	s.clOrdId = &clOrdId
	return s
}

func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/cancel-order",
		secType:  secTypeSigned,
	}

	r.setBodyParam("instId", s.instId)

	if s.ordId != nil {
		r.setBodyParam("ordId", *s.ordId)
	}
	if s.clOrdId != nil {
		r.setBodyParam("clOrdId", *s.clOrdId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CancelOrderServiceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ClOrdId string `json:"clOrdId"`
		OrdId   string `json:"ordId"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}

// AmendOrderService
type AmendOrderService struct {
	c       *Client
	instId  string
	ordId   *string
	clOrdId *string
	newSz   *string
	newPx   *string
}

func (s *AmendOrderService) InstrumentId(instId string) *AmendOrderService {
	s.instId = instId
	return s
}

func (s *AmendOrderService) OrderId(ordId string) *AmendOrderService {
	s.ordId = &ordId
	return s
}

func (s *AmendOrderService) ClientOrderId(clOrdId string) *AmendOrderService {
	s.clOrdId = &clOrdId
	return s
}

func (s *AmendOrderService) NewSize(newSz string) *AmendOrderService {
	s.newSz = &newSz
	return s
}

func (s *AmendOrderService) NewPrice(newPx string) *AmendOrderService {
	s.newPx = &newPx
	return s
}

func (s *AmendOrderService) Do(ctx context.Context, opts ...RequestOption) (res *EditOrderServiceResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/api/v5/trade/amend-order",
		secType:  secTypeSigned,
	}

	r.setBodyParam("instId", s.instId)

	if s.ordId != nil {
		r.setBodyParam("ordId", *s.ordId)
	}
	if s.clOrdId != nil {
		r.setBodyParam("clOrdId", *s.clOrdId)
	}
	if s.newSz != nil {
		r.setBodyParam("newSz", *s.newSz)
	}
	if s.newPx != nil {
		r.setBodyParam("newPx", *s.newPx)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(EditOrderServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type EditOrderServiceResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ClOrdId string `json:"clOrdId"`
		OrdId   string `json:"ordId"`
		Tag     string `json:"tag"`
		SCode   string `json:"sCode"`
		SMsg    string `json:"sMsg"`
	} `json:"data"`
}
