package market

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/uncle-gua/okx"
)

type (
	Ticker struct {
		InstID    string             `json:"instId"`
		Last      float64            `json:"last,string"`
		LastSz    float64            `json:"lastSz,string"`
		AskPx     float64            `json:"askPx,string"`
		AskSz     float64            `json:"askSz,string"`
		BidPx     float64            `json:"bidPx,string"`
		BidSz     float64            `json:"bidSz,string"`
		Open24h   float64            `json:"open24h,string"`
		High24h   float64            `json:"high24h,string"`
		Low24h    float64            `json:"low24h,string"`
		VolCcy24h float64            `json:"volCcy24h,string"`
		Vol24h    float64            `json:"vol24h,string"`
		SodUtc0   float64            `json:"sodUtc0,string"`
		SodUtc8   float64            `json:"sodUtc8,string"`
		InstType  okx.InstrumentType `json:"instType"`
		TS        int64              `json:"ts,string"`
	}
	IndexTicker struct {
		InstID  string  `json:"instId"`
		IdxPx   float64 `json:"idxPx,string"`
		High24h float64 `json:"high24h,string"`
		Low24h  float64 `json:"low24h,string"`
		Open24h float64 `json:"open24h,string"`
		SodUtc0 float64 `json:"sodUtc0,string"`
		SodUtc8 float64 `json:"sodUtc8,string"`
		TS      int64   `json:"ts,string"`
	}
	OrderBook struct {
		Asks []OrderBookEntity `json:"asks"`
		Bids []OrderBookEntity `json:"bids"`
		TS   int64             `json:"ts,string"`
	}
	OrderBookWs struct {
		Asks     []OrderBookEntity `json:"asks"`
		Bids     []OrderBookEntity `json:"bids"`
		Checksum int               `json:"checksum"`
		TS       int64             `json:"ts,string"`
	}
	OrderBookEntity struct {
		DepthPrice      float64
		Size            float64
		LiquidatedOrder int
		OrderNumbers    int
	}
	Candle struct {
		O      float64
		H      float64
		L      float64
		C      float64
		Vol    float64
		VolCcy float64
		TS     int64
	}
	IndexCandle struct {
		O  float64
		H  float64
		L  float64
		C  float64
		TS int64
	}
	Trade struct {
		InstID  string        `json:"instId"`
		TradeID float64       `json:"tradeId,string"`
		Px      float64       `json:"px,string"`
		Sz      float64       `json:"sz,string"`
		Side    okx.TradeSide `json:"side"`
		TS      int64         `json:"ts,string"`
	}
	TotalVolume24H struct {
		VolUsd float64 `json:"volUsd,string"`
		VolCny float64 `json:"volCny,string"`
		TS     int64   `json:"ts,string"`
	}
	IndexComponent struct {
		Index      string      `json:"index"`
		Last       float64     `json:"last,string"`
		Components []Component `json:"components"`
		TS         int64       `json:"ts,string"`
	}
	Component struct {
		Exch   string  `json:"exch"`
		Symbol string  `json:"symbol"`
		SymPx  float64 `json:"symPx,string"`
		Wgt    float64 `json:"wgt,string"`
		CnvPx  float64 `json:"cnvPx,string"`
	}
)

func (o *OrderBookEntity) UnmarshalJSON(buf []byte) error {
	var (
		dp, s, lo, on string
		err           error
	)
	tmp := []interface{}{&dp, &s, &lo, &on}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in OrderBookEntity: %d != %d", g, e)
	}
	o.DepthPrice, err = strconv.ParseFloat(dp, 64)
	if err != nil {
		return err
	}
	o.Size, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	o.LiquidatedOrder, err = strconv.Atoi(lo)
	if err != nil {
		return err
	}
	o.OrderNumbers, err = strconv.Atoi(on)
	if err != nil {
		return err
	}

	return nil
}

func (c *Candle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, vol, volCcy, ts string
		err                          error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl, &vol, &volCcy}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	c.TS = timestamp

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	c.Vol, err = strconv.ParseFloat(vol, 64)
	if err != nil {
		return err
	}

	c.VolCcy, err = strconv.ParseFloat(volCcy, 64)
	if err != nil {
		return err
	}

	return nil
}

func (c *IndexCandle) UnmarshalJSON(buf []byte) error {
	var (
		o, h, l, cl, ts string
		err             error
	)
	tmp := []interface{}{&ts, &o, &h, &l, &cl}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}

	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Candle: %d != %d", g, e)
	}

	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	c.TS = timestamp

	c.O, err = strconv.ParseFloat(o, 64)
	if err != nil {
		return err
	}

	c.H, err = strconv.ParseFloat(h, 64)
	if err != nil {
		return err
	}

	c.L, err = strconv.ParseFloat(l, 64)
	if err != nil {
		return err
	}

	c.C, err = strconv.ParseFloat(cl, 64)
	if err != nil {
		return err
	}

	return nil
}
