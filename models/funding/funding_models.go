package funding

import "github.com/uncle-gua/okx"

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string          `json:"transId"`
		Ccy     string          `json:"ccy"`
		Amt     float64         `json:"amt,string"`
		From    okx.AccountType `json:"from,string"`
		To      okx.AccountType `json:"to,string"`
	}
	Bill struct {
		BillID string       `json:"billId"`
		Ccy    string       `json:"ccy"`
		Bal    float64      `json:"bal,string"`
		BalChg float64      `json:"balChg,string"`
		Type   okx.BillType `json:"type,string"`
		TS     int64        `json:"ts,string"`
	}
	DepositAddress struct {
		Addr     string          `json:"addr"`
		Tag      string          `json:"tag,omitempty"`
		Memo     string          `json:"memo,omitempty"`
		PmtID    string          `json:"pmtId,omitempty"`
		Ccy      string          `json:"ccy"`
		Chain    string          `json:"chain"`
		CtAddr   string          `json:"ctAddr"`
		Selected bool            `json:"selected"`
		To       okx.AccountType `json:"to,string"`
		TS       int64           `json:"ts,string"`
	}
	DepositHistory struct {
		Ccy   string           `json:"ccy"`
		Chain string           `json:"chain"`
		TxID  string           `json:"txId"`
		From  string           `json:"from"`
		To    string           `json:"to"`
		DepId string           `json:"depId"`
		Amt   float64          `json:"amt,string"`
		State okx.DepositState `json:"state,string"`
		TS    int64            `json:"ts,string"`
	}
	Withdrawal struct {
		Ccy   string  `json:"ccy"`
		Chain string  `json:"chain"`
		WdID  int64   `json:"wdId,string"`
		Amt   float64 `json:"amt,string"`
	}
	WithdrawalHistory struct {
		Ccy   string              `json:"ccy"`
		Chain string              `json:"chain"`
		TxID  string              `json:"txId"`
		From  string              `json:"from"`
		To    string              `json:"to"`
		Tag   string              `json:"tag,omitempty"`
		PmtID string              `json:"pmtId,omitempty"`
		Memo  string              `json:"memo,omitempty"`
		Amt   float64             `json:"amt,string"`
		Fee   float64             `json:"fee,string"`
		WdID  int64               `json:"wdId,string"`
		State okx.WithdrawalState `json:"state,string"`
		TS    int64               `json:"ts,string"`
	}
	PiggyBank struct {
		Ccy  string         `json:"ccy"`
		Amt  float64        `json:"amt,string"`
		Side okx.ActionType `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string  `json:"ccy"`
		Amt      float64 `json:"amt,string"`
		Earnings float64 `json:"earnings,string"`
	}
)
