package trade

import "github.com/uncle-gua/okx"

type (
	PlaceOrder struct {
		ClOrdID string  `json:"clOrdId"`
		Tag     string  `json:"tag"`
		SMsg    string  `json:"sMsg"`
		SCode   int64   `json:"sCode,string"`
		OrdID   float64 `json:"ordId,string"`
	}
	CancelOrder struct {
		OrdID   string  `json:"ordId"`
		ClOrdID string  `json:"clOrdId"`
		SMsg    string  `json:"sMsg"`
		SCode   float64 `json:"sCode,string"`
	}
	AmendOrder struct {
		OrdID   string  `json:"ordId"`
		ClOrdID string  `json:"clOrdId"`
		ReqID   string  `json:"reqId"`
		SMsg    string  `json:"sMsg"`
		SCode   float64 `json:"sCode,string"`
	}
	ClosePosition struct {
		InstID  string           `json:"instId"`
		PosSide okx.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string             `json:"instId"`
		Ccy         string             `json:"ccy"`
		OrdID       string             `json:"ordId"`
		ClOrdID     string             `json:"clOrdId"`
		TradeID     string             `json:"tradeId"`
		Tag         string             `json:"tag"`
		Category    string             `json:"category"`
		FeeCcy      string             `json:"feeCcy"`
		RebateCcy   string             `json:"rebateCcy"`
		Px          float64            `json:"px,string"`
		Sz          float64            `json:"sz,string"`
		Pnl         float64            `json:"pnl,string"`
		AccFillSz   float64            `json:"accFillSz,string"`
		FillPx      float64            `json:"fillPx,string"`
		FillSz      float64            `json:"fillSz,string"`
		FillTime    float64            `json:"fillTime,string"`
		AvgPx       float64            `json:"avgPx,string"`
		Lever       float64            `json:"lever,string"`
		TpTriggerPx float64            `json:"tpTriggerPx,string"`
		TpOrdPx     float64            `json:"tpOrdPx,string"`
		SlTriggerPx float64            `json:"slTriggerPx,string"`
		SlOrdPx     float64            `json:"slOrdPx,string"`
		Fee         float64            `json:"fee,string"`
		Rebate      float64            `json:"rebate,string"`
		State       okx.OrderState     `json:"state"`
		TdMode      okx.TradeMode      `json:"tdMode"`
		PosSide     okx.PositionSide   `json:"posSide"`
		Side        okx.OrderSide      `json:"side"`
		OrdType     okx.OrderType      `json:"ordType"`
		InstType    okx.InstrumentType `json:"instType"`
		TgtCcy      okx.QuantityType   `json:"tgtCcy"`
		UTime       int64              `json:"uTime,string"`
		CTime       int64              `json:"cTime,string"`
	}
	TransactionDetail struct {
		InstID   string             `json:"instId"`
		OrdID    string             `json:"ordId"`
		TradeID  string             `json:"tradeId"`
		ClOrdID  string             `json:"clOrdId"`
		BillID   string             `json:"billId"`
		Tag      float64            `json:"tag,string"`
		FillPx   float64            `json:"fillPx,string"`
		FillSz   float64            `json:"fillSz,string"`
		FeeCcy   float64            `json:"feeCcy,string"`
		Fee      float64            `json:"fee,string"`
		InstType okx.InstrumentType `json:"instType"`
		Side     okx.OrderSide      `json:"side"`
		PosSide  okx.PositionSide   `json:"posSide"`
		ExecType okx.OrderFlowType  `json:"execType"`
		TS       int64              `json:"ts,string"`
	}
	PlaceAlgoOrder struct {
		AlgoID string `json:"algoId"`
		SMsg   string `json:"sMsg"`
		SCode  int64  `json:"sCode,string"`
	}
	CancelAlgoOrder struct {
		AlgoID string `json:"algoId"`
		SMsg   string `json:"sMsg"`
		SCode  int64  `json:"sCode,string"`
	}
	AlgoOrder struct {
		InstID       string             `json:"instId"`
		Ccy          string             `json:"ccy"`
		OrdID        string             `json:"ordId"`
		AlgoID       string             `json:"algoId"`
		ClOrdID      string             `json:"clOrdId"`
		TradeID      string             `json:"tradeId"`
		Tag          string             `json:"tag"`
		Category     string             `json:"category"`
		FeeCcy       string             `json:"feeCcy"`
		RebateCcy    string             `json:"rebateCcy"`
		TimeInterval string             `json:"timeInterval"`
		Px           float64            `json:"px,string"`
		PxVar        float64            `json:"pxVar,string"`
		PxSpread     float64            `json:"pxSpread,string"`
		PxLimit      float64            `json:"pxLimit,string"`
		Sz           float64            `json:"sz,string"`
		SzLimit      float64            `json:"szLimit,string"`
		ActualSz     float64            `json:"actualSz,string"`
		ActualPx     float64            `json:"actualPx,string"`
		Pnl          float64            `json:"pnl,string"`
		AccFillSz    float64            `json:"accFillSz,string"`
		FillPx       float64            `json:"fillPx,string"`
		FillSz       float64            `json:"fillSz,string"`
		FillTime     float64            `json:"fillTime,string"`
		AvgPx        float64            `json:"avgPx,string"`
		Lever        float64            `json:"lever,string"`
		TpTriggerPx  float64            `json:"tpTriggerPx,string"`
		TpOrdPx      float64            `json:"tpOrdPx,string"`
		SlTriggerPx  float64            `json:"slTriggerPx,string"`
		SlOrdPx      float64            `json:"slOrdPx,string"`
		OrdPx        float64            `json:"ordPx,string"`
		Fee          float64            `json:"fee,string"`
		Rebate       float64            `json:"rebate,string"`
		State        okx.OrderState     `json:"state"`
		TdMode       okx.TradeMode      `json:"tdMode"`
		ActualSide   okx.PositionSide   `json:"actualSide"`
		PosSide      okx.PositionSide   `json:"posSide"`
		Side         okx.OrderSide      `json:"side"`
		OrdType      okx.AlgoOrderType  `json:"ordType"`
		InstType     okx.InstrumentType `json:"instType"`
		TgtCcy       okx.QuantityType   `json:"tgtCcy"`
		CTime        int64              `json:"cTime,string"`
		TriggerTime  int64              `json:"triggerTime,string"`
	}
)
