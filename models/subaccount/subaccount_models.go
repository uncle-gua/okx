package subaccount

import "github.com/uncle-gua/okx"

type (
	SubAccount struct {
		SubAcct string `json:"subAcct,omitempty"`
		Label   string `json:"label,omitempty"`
		Mobile  string `json:"mobile,omitempty"`
		GAuth   bool   `json:"gAuth"`
		Enable  bool   `json:"enable"`
		TS      int64  `json:"ts,string"`
	}
	APIKey struct {
		SubAcct    string `json:"subAcct,omitempty"`
		Label      string `json:"label,omitempty"`
		APIKey     string `json:"apiKey,omitempty"`
		SecretKey  string `json:"secretKey,omitempty"`
		Passphrase string `json:"Passphrase,omitempty"`
		Perm       string `json:"perm,omitempty"`
		IP         string `json:"ip,omitempty"`
		TS         int64  `json:"ts,omitempty,string"`
	}
	HistoryTransfer struct {
		SubAcct string       `json:"subAcct,omitempty"`
		Ccy     string       `json:"ccy,omitempty"`
		BillID  int64        `json:"billId,omitempty,string"`
		Type    okx.BillType `json:"type,omitempty,string"`
		TS      int64        `json:"ts,omitempty,string"`
	}
	Transfer struct {
		TransID int64 `json:"transId,string"`
	}
)
