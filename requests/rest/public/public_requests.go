package public

import "github.com/uncle-gua/okx"

type (
	GetInstruments struct {
		Uly      string             `json:"uly,omitempty"`
		InstID   string             `json:"instId,omitempty"`
		InstType okx.InstrumentType `json:"instType"`
	}
	GetDeliveryExerciseHistory struct {
		Uly      string             `json:"uly"`
		After    int64              `json:"after,omitempty,string"`
		Before   int64              `json:"before,omitempty,string"`
		Limit    int64              `json:"limit,omitempty,string"`
		InstType okx.InstrumentType `json:"instType"`
	}
	GetOpenInterest struct {
		Uly      string             `json:"uly,omitempty"`
		InstID   string             `json:"instId,omitempty"`
		InstType okx.InstrumentType `json:"instType"`
	}
	GetFundingRate struct {
		InstID string `json:"instId"`
	}
	GetLimitPrice struct {
		InstID string `json:"instId"`
	}
	GetOptionMarketData struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetEstimatedDeliveryExercisePrice struct {
		Uly     string `json:"uly"`
		ExpTime string `json:"expTime,omitempty"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Uly        string  `json:"uly"`
		Ccy        string  `json:"ccy,omitempty"`
		DiscountLv float64 `json:"discountLv,string"`
	}
	GetLiquidationOrders struct {
		InstID   string             `json:"instId,omitempty"`
		Ccy      string             `json:"ccy,omitempty"`
		Uly      string             `json:"uly,omitempty"`
		After    int64              `json:"after,omitempty,string"`
		Before   int64              `json:"before,omitempty,string"`
		Limit    int64              `json:"limit,omitempty,string"`
		InstType okx.InstrumentType `json:"instType"`
		MgnMode  okx.MarginMode     `json:"mgnMode,omitempty"`
		Alias    okx.AliasType      `json:"alias,omitempty"`
		State    okx.OrderState     `json:"state,omitempty"`
	}
	GetMarkPrice struct {
		InstID   string             `json:"instId,omitempty"`
		Uly      string             `json:"uly,omitempty"`
		InstType okx.InstrumentType `json:"instType"`
	}
	GetPositionTiers struct {
		InstID   string             `json:"instId,omitempty"`
		Uly      string             `json:"uly,omitempty"`
		InstType okx.InstrumentType `json:"instType"`
		TdMode   okx.TradeMode      `json:"tdMode"`
		Tier     int64              `json:"tier,omitempty,string"`
	}
	GetUnderlying struct {
		InstType okx.InstrumentType `json:"instType"`
	}
	Status struct {
		State string `json:"state,omitempty"`
	}
)
