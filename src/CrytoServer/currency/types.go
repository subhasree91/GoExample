package currency

type Currency struct {
	ID                 string `json:"id"`
	FullName           string `json:"fullName"`
	Crypto             bool   `json:"crypto"`
	PayinEnabled       bool   `json:"payinEnabled"`
	PayinPaymentID     bool   `json:"payinPaymentId"`
	PayinConfirmations int    `json:"payinConfirmations"`
	PayoutEnabled      bool   `json:"payoutEnabled"`
	PayoutIsPaymentID  bool   `json:"payoutIsPaymentId"`
	TransferEnabled    bool   `json:"transferEnabled"`
	Delisted           bool   `json:"delisted"`
	PayoutFee          string `json:"payoutFee"`
}

type Symbol struct {
	ID                   string `json: "id"`
	BaseCurrency         string `json: "baseCurrency"`
	QuoteCurrency        string `json: "quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json: "tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json: "provideLiquidityRate"`
	FeeCurrency          string `json: "feeCurrency"`
}

const (
	CurrencyAllUrl   = "https://api.hitbtc.com/api/2/public/currency"
	CurencySymbolUrl = "https://api.hitbtc.com/api/2/public/symbol"
)
