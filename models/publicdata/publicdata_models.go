package publicdata

import "github.com/uncle-gua/okx"

type (
	Instrument struct {
		InstID    string              `json:"instId"`
		Uly       string              `json:"uly,omitempty"`
		BaseCcy   string              `json:"baseCcy,omitempty"`
		QuoteCcy  string              `json:"quoteCcy,omitempty"`
		SettleCcy string              `json:"settleCcy,omitempty"`
		CtValCcy  string              `json:"ctValCcy,omitempty"`
		CtVal     float64             `json:"ctVal,omitempty,string"`
		CtMult    float64             `json:"ctMult,omitempty,string"`
		Stk       float64             `json:"stk,omitempty,string"`
		TickSz    float64             `json:"tickSz,omitempty,string"`
		LotSz     float64             `json:"lotSz,omitempty,string"`
		MinSz     float64             `json:"minSz,omitempty,string"`
		Lever     float64             `json:"lever,string"`
		InstType  okx.InstrumentType  `json:"instType"`
		Category  okx.FeeCategory     `json:"category,string"`
		OptType   okx.OptionType      `json:"optType,omitempty"`
		ListTime  int64               `json:"listTime,string"`
		ExpTime   int64               `json:"expTime,omitempty,string"`
		CtType    okx.ContractType    `json:"ctType,omitempty"`
		Alias     okx.AliasType       `json:"alias,omitempty"`
		State     okx.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []DeliveryExerciseHistoryDetails `json:"details"`
		TS      int64                            `json:"ts,string"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                   `json:"instId"`
		Px     float64                  `json:"px,string"`
		Type   okx.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string             `json:"instId"`
		Oi       float64            `json:"oi,string"`
		OiCcy    float64            `json:"oiCcy,string"`
		InstType okx.InstrumentType `json:"instType"`
		TS       int64              `json:"ts,string"`
	}
	FundingRate struct {
		InstID          string             `json:"instId"`
		InstType        okx.InstrumentType `json:"instType"`
		FundingRate     float64            `json:"fundingRate,string"`
		NextFundingRate float64            `json:"NextFundingRate,string"`
		FundingTime     int64              `json:"fundingTime,string"`
		NextFundingTime int64              `json:"nextFundingTime,string"`
	}
	LimitPrice struct {
		InstID   string             `json:"instId"`
		InstType okx.InstrumentType `json:"instType"`
		BuyLmt   float64            `json:"buyLmt,string"`
		SellLmt  float64            `json:"sellLmt,string"`
		TS       int64              `json:"ts,string"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string             `json:"instId"`
		InstType okx.InstrumentType `json:"instType"`
		SettlePx float64            `json:"settlePx,string"`
		TS       int64              `json:"ts,string"`
	}
	OptionMarketData struct {
		InstID   string             `json:"instId"`
		Uly      string             `json:"uly"`
		InstType okx.InstrumentType `json:"instType"`
		Delta    float64            `json:"delta,string"`
		Gamma    float64            `json:"gamma,string"`
		Vega     float64            `json:"vega,string"`
		Theta    float64            `json:"theta,string"`
		DeltaBS  float64            `json:"deltaBS,string"`
		GammaBS  float64            `json:"gammaBS,string"`
		VegaBS   float64            `json:"vegaBS,string"`
		ThetaBS  float64            `json:"thetaBS,string"`
		Lever    float64            `json:"lever,string"`
		MarkVol  float64            `json:"markVol,string"`
		BidVol   float64            `json:"bidVol,string"`
		AskVol   float64            `json:"askVol,string"`
		RealVol  float64            `json:"realVol,string"`
		TS       int64              `json:"ts,string"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string         `json:"ccy"`
		Amt          float64        `json:"amt,string"`
		DiscountLv   int64          `json:"discountLv,string"`
		DiscountInfo []DiscountInfo `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate int64 `json:"discountRate,string"`
		MaxAmt       int64 `json:"maxAmt,string"`
		MinAmt       int64 `json:"minAmt,string"`
	}
	SystemTime struct {
		TS int64 `json:"ts,string"`
	}
	LiquidationOrder struct {
		InstID    string                   `json:"instId"`
		Uly       string                   `json:"uly,omitempty"`
		InstType  okx.InstrumentType       `json:"instType"`
		TotalLoss float64                  `json:"totalLoss,string"`
		Details   []LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string           `json:"ccy,omitempty"`
		Side    okx.OrderSide    `json:"side"`
		OosSide okx.PositionSide `json:"posSide"`
		BkPx    float64          `json:"bkPx,string"`
		Sz      float64          `json:"sz,string"`
		BkLoss  float64          `json:"bkLoss,string"`
		TS      int64            `json:"ts,string"`
	}
	MarkPrice struct {
		InstID   string             `json:"instId"`
		InstType okx.InstrumentType `json:"instType"`
		MarkPx   float64            `json:"markPx,string"`
		TS       int64              `json:"ts,string"`
	}
	PositionTier struct {
		InstID       string             `json:"instId"`
		Uly          string             `json:"uly,omitempty"`
		InstType     okx.InstrumentType `json:"instType"`
		Tier         int64              `json:"tier,string"`
		MinSz        float64            `json:"minSz,string"`
		MaxSz        float64            `json:"maxSz,string"`
		Mmr          float64            `json:"mmr,string"`
		Imr          float64            `json:"imr,string"`
		OptMgnFactor float64            `json:"optMgnFactor,omitempty,string"`
		QuoteMaxLoan float64            `json:"quoteMaxLoan,omitempty,string"`
		BaseMaxLoan  float64            `json:"baseMaxLoan,omitempty,string"`
		MaxLever     float64            `json:"maxLever,string"`
		TS           int64              `json:"ts,string"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []InterestRateAndLoanBasic `json:"basic"`
		Vip     []InterestRateAndLoanUser  `json:"vip"`
		Regular []InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string  `json:"ccy"`
		Rate  float64 `json:"rate,string"`
		Quota float64 `json:"quota,string"`
	}
	InterestRateAndLoanUser struct {
		Level         string  `json:"level"`
		IrDiscount    float64 `json:"irDiscount,string"`
		LoanQuotaCoef int     `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string `json:"title"`
		State       string `json:"state"`
		Href        string `json:"href"`
		ServiceType string `json:"serviceType"`
		System      string `json:"system"`
		ScheDesc    string `json:"scheDesc"`
		Begin       int64  `json:"begin,string"`
		End         int64  `json:"end,string"`
	}
)
