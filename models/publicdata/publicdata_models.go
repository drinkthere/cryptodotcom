package publicdata

import "github.com/drinkthere/cryptodotcom"

type (
	Instrument struct {
		Symbol            string                      `json:"symbol"`
		InstType          cryptodotcom.InstrumentType `json:"inst_type"`
		DisplayName       string                      `json:"display_name,,omitempty"`
		BaseCcy           string                      `json:"base_ccy,omitempty"`
		QuoteCcy          string                      `json:"quote_ccy,omitempty"`
		QuoteDecimals     int                         `json:"quote_decimals,omitempty"`
		QuantityDecimals  int                         `json:"quantity_decimals,omitempty"`
		PriceTickSize     string                      `json:"price_tick_size,omitempty"`
		QtyTickSize       string                      `json:"qty_tick_size,omitempty"`
		MaxLeverage       string                      `json:"max_leverage,omitempty"`
		Tradable          bool                        `json:"tradable,omitempty"`
		ExpiryTimestampMs int                         `json:"expiry_timestamp_ms,omitempty"`
		UnderlyingSymbol  string                      `json:"underlying_symbol,omitempty"`
	}
)
