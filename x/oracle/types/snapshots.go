package types

// OracleExchangeRates - array of OracleExchangeRate
type PriceSnapshots []PriceSnapshot

type (
	PriceSnapshotItems []PriceSnapshotItem
	OracleTwaps        []OracleTwap
)

// String implements fmt.Stringer interface
func (snapshots PriceSnapshots) String() string { _ = "STUB: not implemented"; return "" }

// String implements fmt.Stringer interface
func (items PriceSnapshotItems) String() string { _ = "STUB: not implemented"; return "" }

func NewPriceSnapshotItem(denom string, exchangeRate OracleExchangeRate) PriceSnapshotItem {
	_ = "STUB: not implemented"
	return *new(PriceSnapshotItem)
}

func NewPriceSnapshot(priceSnapshotItems PriceSnapshotItems, snapshotTimestamp int64) PriceSnapshot {
	_ = "STUB: not implemented"
	return *new(PriceSnapshot)
}
