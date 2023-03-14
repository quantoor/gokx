package okx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// TradeMode define trade mode
type TradeMode string

// SideType define side type of orders
type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// UserDataEventType define spot user data event type
type UserDataEventType string

// MarginTransferType define margin transfer type
type MarginTransferType int

// MarginLoanStatusType define margin loan status type
type MarginLoanStatusType string

// MarginRepayStatusType define margin repay status type
type MarginRepayStatusType string

// FuturesTransferStatusType define futures transfer status type
type FuturesTransferStatusType string

// SideEffectType define side effect type for orders
type SideEffectType string

// FuturesTransferType define futures transfer type
type FuturesTransferType int

// Endpoints
const (
	baseAPIMainURL = "https://www.okx.com"
)

// UseTestnet switch all the API endpoints from production to the testnet
var UseTestnet = false

// Global enums
const (
	TradeModeIsolated TradeMode = "isolated"
	TradeModeCross    TradeMode = "cross"
	TradeModeCash     TradeMode = "cash"

	SideTypeBuy  SideType = "buy"
	SideTypeSell SideType = "sell"

	PositionSideTypeNet   PositionSideType = "net"
	PositionSideTypeLong  PositionSideType = "long"
	PositionSideTypeShort PositionSideType = "short"

	// Order type
	OrderTypeLimit           OrderType = "limit"
	OrderTypeMarket          OrderType = "market"
	OrderTypePostOnly        OrderType = "post_only"
	OrderTypeFOK             OrderType = "fok"
	OrderTypeIOC             OrderType = "ioc"
	OrderTypeOptimalLimitIOC OrderType = "optimal_limit_ioc"

	// Algo order type
	OrderTypeConditional OrderType = "conditional"
	OrderTypeOCO         OrderType = "oco"
	OrderTypeTrigger     OrderType = "trigger"
	OrderTypeIceberg     OrderType = "iceberg"
	OrderTypeTwap        OrderType = "twap"

	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint() string {

	return baseAPIMainURL
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey, passPhrase string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		PassPhrase: passPhrase,
		BaseURL:    getAPIEndpoint(),
		UserAgent:  "Okx/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Okx-golang ", log.LstdFlags),
		Debug:      false,
		Simulated:  false, // True to enable simulated mode
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	PassPhrase string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Simulated  bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	timestamp := IsoTime()
	queryString := r.query.Encode()
	body := &bytes.Buffer{}

	bodyJson := r.bodyJson
	header := http.Header{}

	if r.header != nil {
		header = r.header.Clone()
	}

	if bodyJson != nil || r.body != nil {
		header.Set("Content-Type", "application/json")
		postBody, _ := json.Marshal(bodyJson)
		body = bytes.NewBuffer(postBody)
	}

	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("OK-ACCESS-KEY", c.APIKey)
		header.Set("OK-ACCESS-PASSPHRASE", c.PassPhrase)
		header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	}

	if c.Simulated {
		header.Set("x-simulated-trading", "1")
	}

	path := r.endpoint
	if queryString != "" {
		path = fmt.Sprintf("%s?%s&", r.endpoint, queryString)
	}
	c.debug("path:" + path)

	if r.body != nil {
		body.Reset()
		body.ReadFrom(r.body)
	}
	if r.secType == secTypeSigned {
		sign, err := Hmac256(timestamp, r.method, path, body, c.SecretKey)
		if err != nil {
			return err
		}

		header.Set("OK-ACCESS-SIGN", sign)
		v := url.Values{}

		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	r.fullURL = fullURL
	r.header = header
	r.body = body

	c.debug("full url: %s, body: %s", fullURL, r.body)
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

// Trade
func (c *Client) NewPlaceOrderService() *PlaceOrderService {
	return &PlaceOrderService{c: c}
}

func (c *Client) NewCancelOrderService() *CancelOrderService {
	return &CancelOrderService{c: c}
}

func (c *Client) NewCancelMultipleOrdersService() *CancelMultipleOrdersService {
	return &CancelMultipleOrdersService{c: c}
}

func (c *Client) NewAmendOrderService() *AmendOrderService {
	return &AmendOrderService{c: c}
}

func (c *Client) NewGetOrderListService() *GetOrderListService {
	return &GetOrderListService{c: c}
}

// Market Data
func (c *Client) NewGetTickerService() *GetTickerService {
	return &GetTickerService{c: c}
}

// Public Data
func (c *Client) NewGetInstrumentsService() *GetInstrumentsService {
	return &GetInstrumentsService{c: c}
}

func (c *Client) NewGetLimitPriceService() *GetLimitPriceService {
	return &GetLimitPriceService{c: c}
}

func (c *Client) SubscribeOrderEvents(eventHandler func(event *WsOrdersEvent), errorHandler func(error)) error {
	go func() {
		_, _, err := WsOrdersServe("ANY", "", "", c.APIKey, c.SecretKey, c.PassPhrase, eventHandler, errorHandler, false)
		if err != nil {
			fmt.Println("ERROR", err)
		}
	}()

	return nil
}
