package account

import "github.com/uncle-gua/okx"

type (
	Balance struct {
		TotalEq     float64          `json:"totalEq,string"`
		IsoEq       float64          `json:"isoEq,string"`
		AdjEq       float64          `json:"adjEq,omitempty,string"`
		OrdFroz     float64          `json:"ordFroz,omitempty,string"`
		Imr         float64          `json:"imr,omitempty,string"`
		Mmr         float64          `json:"mmr,omitempty,string"`
		MgnRatio    float64          `json:"mgnRatio,omitempty,string"`
		NotionalUsd float64          `json:"notionalUsd,omitempty,string"`
		Details     []BalanceDetails `json:"details,omitempty"`
		UTime       int64            `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string  `json:"ccy"`
		Eq            float64 `json:"eq,string"`
		CashBal       float64 `json:"cashBal,string"`
		IsoEq         float64 `json:"isoEq,omitempty,string"`
		AvailEq       float64 `json:"availEq,omitempty,string"`
		DisEq         float64 `json:"disEq,string"`
		AvailBal      float64 `json:"availBal,string"`
		FrozenBal     float64 `json:"frozenBal,string"`
		OrdFrozen     float64 `json:"ordFrozen,string"`
		Liab          float64 `json:"liab,omitempty,string"`
		Upl           float64 `json:"upl,omitempty,string"`
		UplLib        float64 `json:"uplLib,omitempty,string"`
		CrossLiab     float64 `json:"crossLiab,omitempty,string"`
		IsoLiab       float64 `json:"isoLiab,omitempty,string"`
		MgnRatio      float64 `json:"mgnRatio,omitempty,string"`
		Interest      float64 `json:"interest,omitempty,string"`
		Twap          float64 `json:"twap,omitempty,string"`
		MaxLoan       float64 `json:"maxLoan,omitempty,string"`
		EqUsd         float64 `json:"eqUsd,string"`
		NotionalLever float64 `json:"notionalLever,omitempty,string"`
		StgyEq        float64 `json:"stgyEq,string"`
		IsoUpl        float64 `json:"isoUpl,omitempty,string"`
		UTime         int64   `json:"uTime,string"`
	}
	Position struct {
		InstID      string             `json:"instId"`
		PosCcy      string             `json:"posCcy,omitempty"`
		LiabCcy     string             `json:"liabCcy,omitempty"`
		OptVal      string             `json:"optVal,omitempty"`
		Ccy         string             `json:"ccy"`
		PosID       string             `json:"posId"`
		TradeID     string             `json:"tradeId"`
		Pos         float64            `json:"pos,string"`
		AvailPos    float64            `json:"availPos,omitempty,string"`
		AvgPx       float64            `json:"avgPx,string"`
		Upl         float64            `json:"upl,string"`
		UplRatio    float64            `json:"uplRatio,string"`
		Lever       float64            `json:"lever,string"`
		LiqPx       float64            `json:"liqPx,omitempty,string"`
		Imr         float64            `json:"imr,omitempty,string"`
		Margin      float64            `json:"margin,omitempty,string"`
		MgnRatio    float64            `json:"mgnRatio,string"`
		Mmr         float64            `json:"mmr,string"`
		Liab        float64            `json:"liab,omitempty,string"`
		Interest    float64            `json:"interest,string"`
		NotionalUsd float64            `json:"notionalUsd,string"`
		ADL         float64            `json:"adl,string"`
		Last        float64            `json:"last,string"`
		DeltaBS     float64            `json:"deltaBS,string"`
		DeltaPA     float64            `json:"deltaPA,string"`
		GammaBS     float64            `json:"gammaBS,string"`
		GammaPA     float64            `json:"gammaPA,string"`
		ThetaBS     float64            `json:"thetaBS,string"`
		ThetaPA     float64            `json:"thetaPA,string"`
		VegaBS      float64            `json:"vegaBS,string"`
		VegaPA      float64            `json:"vegaPA,string"`
		PosSide     okx.PositionSide   `json:"posSide"`
		MgnMode     okx.MarginMode     `json:"mgnMode"`
		InstType    okx.InstrumentType `json:"instType"`
		CTime       int64              `json:"cTime,string"`
		UTime       int64              `json:"uTime,string"`
	}
	BalanceAndPosition struct {
		EventType okx.EventType    `json:"eventType"`
		PTime     int64            `json:"pTime,string"`
		UTime     int64            `json:"uTime,string"`
		PosData   []Position       `json:"posData"`
		BalData   []BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   float64                             `json:"adjEq,omitempty,string"`
		BalData []PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []PositionAndAccountRiskBalanceData `json:"posData"`
		TS      int64                               `json:"ts,string"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string  `json:"ccy"`
		Eq    float64 `json:"eq,string"`
		DisEq float64 `json:"disEq,string"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string             `json:"instId"`
		PosCcy      string             `json:"posCcy,omitempty"`
		Ccy         string             `json:"ccy"`
		NotionalCcy float64            `json:"notionalCcy,string"`
		Pos         float64            `json:"pos,string"`
		NotionalUsd float64            `json:"notionalUsd,string"`
		PosSide     okx.PositionSide   `json:"posSide"`
		InstType    okx.InstrumentType `json:"instType"`
		MgnMode     okx.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string             `json:"ccy"`
		InstID    string             `json:"instId"`
		Notes     string             `json:"notes"`
		BillID    string             `json:"billId"`
		OrdID     string             `json:"ordId"`
		BalChg    float64            `json:"balChg,string"`
		PosBalChg float64            `json:"posBalChg,string"`
		Bal       float64            `json:"bal,string"`
		PosBal    float64            `json:"posBal,string"`
		Sz        float64            `json:"sz,string"`
		Pnl       float64            `json:"pnl,string"`
		Fee       float64            `json:"fee,string"`
		From      okx.AccountType    `json:"from,string"`
		To        okx.AccountType    `json:"to,string"`
		InstType  okx.InstrumentType `json:"instType"`
		MgnMode   okx.MarginMode     `json:"MgnMode"`
		Type      okx.BillType       `json:"type,string"`
		SubType   okx.BillSubType    `json:"subType,string"`
		TS        int64              `json:"ts,string"`
	}
	Config struct {
		Level      string           `json:"level"`
		LevelTmp   string           `json:"levelTmp"`
		AcctLv     string           `json:"acctLv"`
		AutoLoan   bool             `json:"autoLoan"`
		UID        string           `json:"uid"`
		GreeksType okx.GreekType    `json:"greeksType"`
		PosMode    okx.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okx.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string           `json:"instId"`
		Lever   float64          `json:"lever,string"`
		MgnMode okx.MarginMode   `json:"mgnMode"`
		PosSide okx.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string  `json:"instId"`
		Ccy     string  `json:"ccy"`
		MaxBuy  float64 `json:"maxBuy,string"`
		MaxSell float64 `json:"maxSell,string"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string  `json:"instId"`
		AvailBuy  float64 `json:"availBuy,string"`
		AvailSell float64 `json:"availSell,string"`
	}
	MarginBalanceAmount struct {
		InstID  string           `json:"instId"`
		Amt     float64          `json:"amt,string"`
		PosSide okx.PositionSide `json:"posSide,string"`
		Type    okx.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string         `json:"instId"`
		MgnCcy  string         `json:"mgnCcy"`
		Ccy     string         `json:"ccy"`
		MaxLoan float64        `json:"maxLoan,string"`
		MgnMode okx.MarginMode `json:"mgnMode"`
		Side    okx.OrderSide  `json:"side,string"`
	}
	Fee struct {
		Level    string             `json:"level"`
		Taker    float64            `json:"taker,string"`
		Maker    float64            `json:"maker,string"`
		Delivery float64            `json:"delivery,omitempty,string"`
		Exercise float64            `json:"exercise,omitempty,string"`
		Category okx.FeeCategory    `json:"category,string"`
		InstType okx.InstrumentType `json:"instType"`
		TS       int64              `json:"ts,string"`
	}
	InterestAccrued struct {
		InstID       string         `json:"instId"`
		Ccy          string         `json:"ccy"`
		Interest     float64        `json:"interest,string"`
		InterestRate float64        `json:"interestRate,string"`
		Liab         float64        `json:"liab,string"`
		MgnMode      okx.MarginMode `json:"mgnMode"`
		TS           int64          `json:"ts,string"`
	}
	InterestRate struct {
		Ccy          string  `json:"ccy"`
		InterestRate float64 `json:"interestRate,string"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string  `json:"ccy"`
		MaxWd float64 `json:"maxWd,string"`
	}
)
